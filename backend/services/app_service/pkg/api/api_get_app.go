package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) Get(ctx context.Context, request *appSrv.GetRequest) (*appSrv.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, app, err := srv.app.GetApp(ctx, srv.storage, srv.userService, request.GetAppUuid(), request.GetCallerUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.GetApp] could not GetApp: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.GetResponse{StatusCode: int32(status), Msg: "could not get application data"}, nil
	}
	return &appSrv.GetResponse{
		StatusCode: int32(status),
		Msg:        "app data",
		App:        app,
	}, nil
}
