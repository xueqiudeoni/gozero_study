package user

import (
	"context"
	"fmt"
	"homestay/app/usercenter/cmd/model"
	"homestay/app/usercenter/cmd/rpc/usercenter"
	"homestay/common/xerr"

	"github.com/pkg/errors"
	"github.com/silenceper/wechat/cache"

	"homestay/app/usercenter/cmd/api/internal/svc"
	"homestay/app/usercenter/cmd/api/internal/types"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// wechat mini auth
func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// todo: add your logic here and delete this line
	miniprogram := wechat.NewWechat().GetMiniProgram(&config.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.Appid,
		AppSecret: l.svcCtx.Config.WxMiniConf.Appsecret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID != "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起请求失败，err:%v,code:%s, authResult:%+v", err, req.Code, authResult)
	}
	// 2.parsing WeChat-mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "解析数据失败 req：%+v,err:%v,authResult:%+v", req, err, authResult)
	}
	//bind user or login
	var userId int64
	rpcresp, err := l.svcCtx.UsercenterRpc.GetUserAuthKey(l.ctx, &usercenter.GetUserAuthKeyReq{
		AuthKey:  authResult.OpenID,
		AuthType: model.UserAuthTypeSmallWX,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "rpc call userauthkey err:%v,authResult:%+v", err, authResult)
	}

	if rpcresp.Userrauth == nil || rpcresp.Userrauth.Id == 0 {
		//bind user

		//wechat-mini decrypted data
		mobile := userData.PhoneNumber
		nickname := fmt.Sprintf("homestay%s", mobile[7:])
		registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
			Mobile:   mobile,
			Nickname: nickname,
			AuthKey:  authResult.OpenID,
			AuthType: model.UserAuthTypeSmallWX,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterrpc.register err:%v,authresult:%+v", err, authResult)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  registerResp.AccessToken,
			AccessExpire: registerResp.AccessExpire,
			RefreshAfter: registerResp.AccessAfter,
		}, nil
	} else {
		userId = rpcresp.Userrauth.UserId
		tokenResp, err := l.svcCtx.UsercenterRpc.GenToken(l.ctx, &usercenter.GenTokenReq{UserId: userId})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterrpc.gentoken err:%v,userId:%d", err, userId)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}, nil
	}
}
