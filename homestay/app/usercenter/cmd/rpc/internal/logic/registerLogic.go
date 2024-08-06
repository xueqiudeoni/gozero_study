package logic

import (
	"context"
	"homestay/app/usercenter/cmd/model"
	"homestay/common/tool"
	"homestay/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"homestay/app/usercenter/cmd/rpc/internal/svc"
	"homestay/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterErr = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterErr, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	}

	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = in.Mobile
		if len(in.Nickname) == 0 {
			user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		if len(in.Password) > 0 {
			user.Password = tool.MD5ByString(in.Password)
		}
		insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "register user db insert err:%v,user:%+v", err, user)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "register user last insertResult.LastInsertId err:%v,user:%+v", err, user)
		}
		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "register user auth insert err:%v,userAuth:%+v", err, userAuth)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// Generate the token, so that the service doesn't call rpc internally
	genTokenLogic := NewGenTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := genTokenLogic.GenToken(&usercenter.GenTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrorGenerateToken, "gentoken userId:%d", userId)
	}
	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		AccessAfter:  tokenResp.RefreshAfter,
	}, nil
}
