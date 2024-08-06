package logic

import (
	"context"

	"homestay/app/travel/cmd/rpc/doc/pb"
	"homestay/app/travel/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayDetailLogic) HomestayDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HomestayDetailResp{}, nil
}
