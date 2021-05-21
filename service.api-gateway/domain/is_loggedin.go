package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	authproto "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) IsLoggedIn(ctx context.Context, token string) (*common.UserTokenClaims, errors.ErrApi) {

	resp, err := svc.userauthClient.IsAuthed(ctx, &authproto.IsAuthedRequest{
		Tracing_ID:  ctx_value.GetString(ctx, "tracingID"),
		AccessToken: token,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not authenticate request",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK || !resp.GetIsAuthed() {
		return nil, errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "User is not authenticated",
			Err:    err,
		}
	}
	return resp.GetClaims(), nil
}
