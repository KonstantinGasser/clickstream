package api

import (
	"context"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/jwts"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenService) VerifyUserToken(ctx context.Context, in *tokenSrv.VerifyUserTokenRequest) (*tokenSrv.VerifyUserTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[tokenService.VerifyUserToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	userInfo, err := jwts.GetJWTClaims(ctx, in.GetUserToken())
	if err != nil {
		logrus.Errorf("<%v>[tokenService.VerifyUserToken] could not get JWT claims: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &tokenSrv.VerifyUserTokenResponse{
			StatusCode: http.StatusForbidden,
			Msg:        "user is not authenticated",
			Claim:      nil,
		}, nil
	}
	return &tokenSrv.VerifyUserTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "user is authenticated",
		Claim: &tokenSrv.UserClaim{
			Uuid: userInfo["sub"].(string),
		},
	}, nil
}
