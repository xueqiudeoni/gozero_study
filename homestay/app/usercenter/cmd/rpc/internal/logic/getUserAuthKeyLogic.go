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

type GetUserAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthKeyLogic {
	return &GetUserAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthKeyLogic) GetUserAuthKey(in *user_pb.GetUserAuthKeyReq) (*user_pb.GetUserAuthKeyResp, error) {
	// todo: add your logic here and delete this line
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth fail"), "GetUserAuthKeyLogic FindOneByAuthTypeAuthKey err: %v", err)
	}
	var respUserAuth = &usercenter.UserAuth{}
	_ = copier.Copy(respUserAuth, userAuth)
	return &user_pb.GetUserAuthKeyResp{Userrauth: respUserAuth}, nil
}
