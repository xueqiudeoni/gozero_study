package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"homestay/app/usercenter/cmd/rpc/user_pb"
	"homestay/common/ctxdata"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"homestay/app/usercenter/cmd/rpc/internal/svc"
)

type GenTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenTokenLogic) GenToken(in *user_pb.GenTokenReq) (*user_pb.GenTokenResp, error) {
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.genJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrorGenerateToken, "GenToken userId:%d,err:%v", in.UserId, err)
	}
	return &user_pb.GenTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
func (l *GenTokenLogic) genJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
