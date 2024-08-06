package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"homestay/app/usercenter/cmd/model"
	"homestay/app/usercenter/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlconn, c.CacheRedis),
		UserAuthModel: model.NewUserAuthModel(sqlconn, c.CacheRedis),
	}

}
