package logic

import (
	"context"

	"testdemo/internal/svc"
	"testdemo/testdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditTagLogic {
	return &EditTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditTagLogic) EditTag(in *testdemo.EditTagReq) (*testdemo.EditTagRes, error) {
	// todo: add your logic here and delete this line

	return &testdemo.EditTagRes{}, nil
}
