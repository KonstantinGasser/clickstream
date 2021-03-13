package api

import (
	"errors"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerUserGetByID looks up the user information for a given UUID
func (api API) HandlerUserGetByID(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerUserGetByID] received getByID request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		api.onError(w, errors.New("could not find authenticated user"), http.StatusForbidden)
		return
	}
	resp, err := api.UserSrvClient.GetUserByID(r.Context(), &userSrv.GetUserByIDRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		Uuid:       user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerUserGetByID] could not execute grpc.GetUserByID: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get user information"), http.StatusInternalServerError)
		return
	}
	if resp.GetStatusCode() > 400 {
		logrus.Warn("<%v>[api.HandlerUserGetByID] could not find any user for given UUID\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New(resp.GetMsg()), int(resp.GetStatusCode()))
		return
	}
	api.onScucessJSON(w, map[string]interface{}{
		"status": resp.GetStatusCode(),
		"msg":    resp.GetMsg(),
		"user":   resp.GetUser(),
	}, http.StatusOK)
}
