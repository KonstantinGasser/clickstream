package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppCreate is the api endpoint if a logged in user wants to create a new application
func (api API) HandlerAppCreate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppCreate] received create-app request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	payload, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
	}
	authedUser := ctx_value.GetAuthedUser(r.Context())
	if authedUser == nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not get user-claims, they are nil\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusForbidden)
		return
	}
	respApp, err := api.AppServiceClient.CreateApp(r.Context(), &appSrv.CreateAppRequest{
		OwnerUuid:    authedUser.GetUuid(),
		Name:         payload["app_name"].(string),
		Organization: authedUser.GetOrgnDomain(),
		Tracing_ID:   ctx_value.GetString(r.Context(), "tracingID"),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not execute grpc.CreateApp: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, err, http.StatusInternalServerError)
		return
	}

	api.onScucessJSON(w, map[string]string{
		"msg": respApp.GetMsg(),
	}, int(respApp.GetStatusCode()))
}