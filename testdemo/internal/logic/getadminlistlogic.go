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

type GetAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminListLogic) GetAdminList(in *testdemo.GetAdminListReq) (*testdemo.GetAdminListRes, error) {
	meta := &testdemo.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &testdemo.GetAdminListRes{
		Meta: meta,
	}
	data := &testdemo.GetAdminListData{}
	page := cast.ToInt(in.GetPage())
	pageSize := cast.ToInt(in.GetPageSize())
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	search := make(map[string]interface{})
	search["channelId"] = in.GetChannelId()
	search["name"] = in.GetName()
	admin := models.Admin{}
	util, err := admin.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	total, err := util.GetTotal(l.ctx, &admin, search)
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	data.PageInfo = &testdemo.PageInfo{
		Total:     cast.ToString(total),
		Page:      cast.ToString(page),
		PageSize:  cast.ToString(pageSize),
		TotalPage: util.GetTotalPage(total, pageSize),
	}
	resp.Data = data
	if total <= 0 {
		return resp, nil
	}
	list := make([]models.Admin, 0)
	err = util.GetList(l.ctx, &admin, search, &list, page, pageSize, "id desc")
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	if len(list) <= 0 {
		return resp, nil
	}
	for _, v := range list {
		data.List = append(data.List, v.ReturnData(l.ctx))
	}
	resp.Data = data
	return resp, nil
}
