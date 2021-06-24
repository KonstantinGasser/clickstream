package intercepter

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// TODO: must go in some ENV
type CtxKey string

const (
	secretUser        = "super_secure"
	atuhCtxKey CtxKey = "user"
)

var (
	ErrNotAuthenticated = fmt.Errorf("user is not authenticated")
	ErrJWTParse         = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT       = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
	ErrExpiredJWT       = fmt.Errorf("provided JWT has expired")
)

func WithUnary(middleware grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(middleware)
}

func WithJwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Info("[intercepter.WithJwtAuth] receiced request\n")
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("could not find correct context details")
	}
	accessToken, ok := meta["authorization"]
	if !ok {
		return nil, ErrNotAuthenticated
	}

	authedUser, err := validateJWT(accessToken[0])
	if err != nil {
		return nil, err
	}
	withAuthCtx := context.WithValue(ctx, "user", authedUser)
	withJwtCtx := context.WithValue(withAuthCtx, "authorization", accessToken[0])

	resp, err := handler(withJwtCtx, req)

	return resp, err
}

func validateJWT(accessToken string) (*common.AuthedUser, error) {
	token, err := verifyToken(accessToken, secretUser)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, ErrNotAuthenticated
	}
	if err := claims.Valid(); err != nil {
		return nil, ErrExpiredJWT
	}

	return &common.AuthedUser{
		Uuid:         claims["sub"].(string),
		Organization: claims["orgn"].(string),
		Username:     claims["uname"].(string),
	}, nil
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCorruptJWT
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrJWTParse
	}
	return token, nil
}