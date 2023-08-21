package logic

import (
	"context"
	"github.com/spf13/cast"
	ApiCode "go-common/code"
	"go-common/constant/response"
	"go-common/tools/tracing"
	"testdemo/models"

	"testdemo/internal/svc"
	"testdemo/testdemo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminDetailLogic {
	return &GetAdminDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminDetailLogic) GetAdminDetail(in *testdemo.GetAdminDetailReq) (*testdemo.GetAdminDetailRes, error) {
	meta := &testdemo.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &testdemo.GetAdminDetailRes{
		Meta: meta,
	}
	id := cast.ToUint64(in.GetId())
	if id <= 0 {
		meta.Code = ApiCode.REQ_DATA_ERROR
		meta.Msg = "id不能为空"
		return resp, nil
	}
	admin := models.Admin{}
	util, err := admin.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	err = util.FirstById(l.ctx, &admin, cast.ToUint64(in.GetId()))
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	if !admin.IsExist() {
		meta.Code = ApiCode.REQ_DATA_ERROR
		meta.Msg = "数据不存在"
		return resp, nil
	}
	resp.Data = admin.ReturnData(l.ctx)
	return resp, nil
}
