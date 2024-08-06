package logic

import (
	"context"
	"homestay/app/usercenter/cmd/model"
	"homestay/app/usercenter/cmd/rpc/usercenter"
	"homestay/common/xerr"

	"github.com/pkg/errors"

	"homestay/app/usercenter/cmd/rpc/internal/svc"
	"homestay/app/usercenter/cmd/rpc/user_pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user_pb.GetUserInfoReq) (*user_pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo FindOne err,id:%d,err: %v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNotExist, "id:%d", in.Id)
	}
	respUser := usercenter.User{}

	_ = copier.Copy(respUser, user)
	return &user_pb.GetUserInfoResp{User: &respUser}, nil
}
