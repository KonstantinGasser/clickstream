package api

import (
	"context"
	"fmt"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/jwts"
	"github.com/sirupsen/logrus"
)

func (srv TokenService) IssueJWT(ctx context.Context, request *tokenSrv.IssueJWTRequest) (*tokenSrv.IssueJWTResponse, error) {

	user := request.GetUser()
	token, err := jwts.IssueUser(user.GetUuid(), user.GetUsername(), user.GetOrgnDomain())
	if err != nil {
		logrus.Errorf("[tokenService.IssueJWT] could not issue JWT for user: %v", err)
		return &tokenSrv.IssueJWTResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "could not issue JWT for user",
			JwtToken:   "",
		}, fmt.Errorf("could not issue user JWT: %v", err)
	}
	return &tokenSrv.IssueJWTResponse{
		StatusCode: http.StatusOK,
		Msg:        "JWT for user issued",
		JwtToken:   token,
	}, nil
}
