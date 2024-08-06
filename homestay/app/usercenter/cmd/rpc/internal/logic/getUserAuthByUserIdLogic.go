package logic

import (
	"context"
	"homestay/app/usercenter/cmd/model"
	"homestay/app/usercenter/cmd/rpc/usercenter"
	"homestay/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"homestay/app/usercenter/cmd/rpc/internal/svc"
	"homestay/app/usercenter/cmd/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *user_pb.GetUserAuthByUserIdReq) (*user_pb.GetUserAuthByUserIdResp, error) {
	// todo: add your logic here and delete this line
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get userr auth fail"), "err:%v,in:%+v", err, in)
	}
	var resp usercenter.UserAuth
	copier.Copy(&resp, userAuth)
	return &user_pb.GetUserAuthByUserIdResp{UserAuth: &resp}, nil
}
