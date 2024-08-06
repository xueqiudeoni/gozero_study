package logic

import (
	"context"
	"homestay/app/usercenter/cmd/model"
	"homestay/common/tool"
	"homestay/common/xerr"

	"github.com/pkg/errors"

	"homestay/app/usercenter/cmd/rpc/internal/svc"
	"homestay/app/usercenter/cmd/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrorGenerateToken = xerr.NewErrMsg("生成token失败")
var ErrorUserNamePasswd = xerr.NewErrMsg("账号或密码不正确")
var ErrUserNotExist = xerr.NewErrMsg("用户不存在")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user_pb.LoginReq) (*user_pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	var userId int64
	var err error

	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SEVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}
	genTokenLogic := NewGenTokenLogic(l.ctx, l.svcCtx)
	genTokenResp, err := genTokenLogic.GenToken(&user_pb.GenTokenReq{UserId: userId})
	if err != nil {
		return nil, errors.Wrapf(ErrorGenerateToken, "生成token失败，userId:%d", userId)
	}
	return &user_pb.LoginResp{
		AccessToken:  genTokenResp.AccessToken,
		AccessExpire: genTokenResp.AccessExpire,
		RefreshAfter: genTokenResp.RefreshAfter,
	}, nil
}
func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile：%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNotExist, "mobile:%s", mobile)
	}
	if user.Password != tool.MD5ByString(password) {
		return 0, errors.Wrapf(ErrorUserNamePasswd, "密码不正确")
	}
	return user.Id, nil
}
