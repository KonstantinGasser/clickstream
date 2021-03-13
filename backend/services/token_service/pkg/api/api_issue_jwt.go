package api

import (
	"context"
	"fmt"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/jwts"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// IssueJWT issues a new JWT for a authenticated user only
func (srv TokenService) IssueJWT(ctx context.Context, request *tokenSrv.IssueJWTRequest) (*tokenSrv.IssueJWTResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[tokenService.IssueJWT] received issuing of JWT request\n", ctx_value.GetString(ctx, "tracingID"))
	user := request.GetUser()
	token, err := jwts.IssueUser(ctx, user)
	if err != nil {
		logrus.Errorf("<%v>[tokenService.IssueJWT] could not issue JWT for user: %v", ctx_value.GetString(ctx, "tracingID"), err)
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
