package logic

import (
	"context"

	"testdemo/internal/svc"
	"testdemo/testdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTagLogic {
	return &DelTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTagLogic) DelTag(in *testdemo.DelTagReq) (*testdemo.DelTagRes, error) {
	// todo: add your logic here and delete this line

	return &testdemo.DelTagRes{}, nil
}
