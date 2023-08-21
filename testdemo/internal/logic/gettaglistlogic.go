package logic

import (
	"context"

	"testdemo/internal/svc"
	"testdemo/testdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagListLogic) GetTagList(in *testdemo.GetTagReq) (*testdemo.GetTagRes, error) {
	// todo: add your logic here and delete this line

	return &testdemo.GetTagRes{}, nil
}
