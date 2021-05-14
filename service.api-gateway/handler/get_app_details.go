package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) GetAppDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.GetAppDetails] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		handler.onError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	appUuid := r.URL.Query().Get("app")
	if appUuid == "" {
		handler.onError(w, "App Uuid missing in query", http.StatusBadRequest)
		return
	}
	app, err := handler.domain.GetAppInfo(r.Context(), user.Uuid, appUuid)
	if err != nil {
		logrus.Errorf("<%v>[handler.GetAppDetails] could not get app: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	token, err := handler.domain.GetAppToken(r.Context(), appUuid)
	if err != nil {
		logrus.Warnf("<%v>[handler.GetAppDetails] could not get app token: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
	}
	config, err := handler.domain.GetAppConfig(r.Context(), appUuid)
	if err != nil {
		logrus.Warnf("<%v>[handler.GetAppDetails] could not get app config: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
	}
	owner, err := handler.domain.GetUserProfile(r.Context(), app.Owner)
	if err != nil {
		logrus.Warnf("<%v>[handler.GetAppDetails] could not get app owner: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status": http.StatusOK,
		"app":    app,
		"config": config,
		"token":  token,
		"owner":  owner,
	}, http.StatusOK)
}
