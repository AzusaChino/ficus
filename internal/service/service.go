package service

import (
	"context"

	"github.com/azusachino/ficus/global"
	"github.com/azusachino/ficus/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	s := Service{ctx: ctx}
	s.dao = dao.New(global.DbEngine)
	return s
}
