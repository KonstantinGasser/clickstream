package jwts

import (
	"context"
	"errors"
	"fmt"
	"time"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	jwt "github.com/dgrijalva/jwt-go"
)

// Todo: must go in env
const (
	issuer  = "com.datalabs.token-service"
	secret  = "super_secure"
	expTime = time.Minute * 60
)

// IssueUser takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func IssueUser(ctx context.Context, user *tokenSrv.AuthenticatedUser) (string, error) {
	// calims holds all the data which will be
	// encoded in the JWT
	claims := jwt.MapClaims{}
	claims["uuid"] = user.GetUuid()
	claims["username"] = user.GetUsername()
	claims["orgnDomain"] = user.GetOrgnDomain()
	claims["firstName"] = user.GetFirstName()
	claims["lastName"] = user.GetLastName()
	claims["porifleImgURL"] = user.GetProfileImgUrl()
	claims["orgnPosition"] = user.GetOrgnPosition()
	claims["iat"] = issuer
	claims["exp"] = time.Now().Add(expTime).Unix()

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}
	return token, nil
}

func GetJWTClaims(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("user not authenticated")
	}
	return claims, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("[jwts.verifyToken] could not parse JWT: %v", err)
	}
	return token, nil
}
