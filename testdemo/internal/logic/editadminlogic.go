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

type EditAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditAdminLogic {
	return &EditAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditAdminLogic) EditAdmin(in *testdemo.EditAdminReq) (*testdemo.EditAdminRes, error) {
	meta := &testdemo.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &testdemo.EditAdminRes{
		Meta: meta,
	}
	id := cast.ToUint64(in.GetId())
	admin := models.Admin{}
	util, err := admin.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	if id > 0 {
		err = util.FirstById(l.ctx, &admin, cast.ToUint64(in.GetId()))
		if err != nil {
			meta.Code = ApiCode.DB_ERROR
			meta.Msg = err.Error()
			return resp, nil
		}
		if !admin.IsExist() {
			meta.Code = ApiCode.REQ_DATA_ERROR
			meta.Msg = "不存在该用户"
			return resp, nil
		}
	} else {
		admin.Id = models.GenerateId()
	}
	admin.Name = in.GetName()
	admin.ChannelId = in.GetChannelId()
	err = util.SaveById(l.ctx, &admin, admin.Id)
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	resp.Data = admin.ReturnData(l.ctx)
	return resp, nil
}
