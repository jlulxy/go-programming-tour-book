package service

import (
	"context"

	"github.com/jlulxy/go-programming-tour-book/blog-service/global"

	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
