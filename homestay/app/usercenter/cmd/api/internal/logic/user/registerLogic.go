package user

import (
	"context"
	"homestay/app/usercenter/cmd/model"
	"homestay/app/usercenter/cmd/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"homestay/app/usercenter/cmd/api/internal/svc"
	"homestay/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", req)
	}
	_ = copier.Copy(resp, registerResp)
	return
}
