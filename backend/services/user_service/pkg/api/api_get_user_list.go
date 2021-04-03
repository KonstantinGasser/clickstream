package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) GetUserList(ctx context.Context, in *userSrv.GetUserListRequest) (*userSrv.GetUserListResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.GetUserList] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, userList, err := srv.user.GetByIDs(ctx, srv.storage, in.GetUuidList())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUserList] could not execute GetByIDs: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetUserListResponse{StatusCode: int32(status), Msg: "Could not get users information", UserList: []*userSrv.ComplexUser{}}, nil
	}
	// convert found userList to grpc User slice
	var users []*userSrv.ComplexUser = make([]*userSrv.ComplexUser, len(userList))
	for i, item := range userList {
		users[i] = &userSrv.ComplexUser{
			Uuid:          item.UUID,
			Username:      item.Username,
			FirstName:     item.FirstName,
			LastName:      item.LastName,
			OrgnDomain:    item.OrgnDomain,
			OrgnPosition:  item.OrgnPosition,
			ProfileImgUrl: item.ProfileImgURL,
		}
	}
	return &userSrv.GetUserListResponse{
		StatusCode: int32(status),
		Msg:        "users record by uuids",
		UserList:   users,
	}, nil
}
