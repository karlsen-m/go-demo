package logic

import (
	"context"

	"testdemo/internal/svc"
	"testdemo/testdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagDetailLogic {
	return &GetTagDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagDetailLogic) GetTagDetail(in *testdemo.GetTagReq) (*testdemo.GetTagRes, error) {
	// todo: add your logic here and delete this line

	return &testdemo.GetTagRes{}, nil
}
