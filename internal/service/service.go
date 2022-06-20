package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/yunshu2009/blog-api/global"
	"github.com/yunshu2009/blog-api/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
