package logic

func createEditLogic(pkgName, apiName, modelName, resDataName string) string {
	str := `
	mate := &` + pkgName + `.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &` + pkgName + `.` + apiName + `Res{
		Meta: mate,
	}
	id := cast.ToUint64(in.GetId())
	modelClient := models.NewModelClient("` + modelName + `")
	model := models.` + modelName + `{}
	if id > 0 {
		modelI, err := modelClient.GetById(l.ctx, id)
		if err != nil {
			mate.Code = ApiCode.DB_ERROR
			mate.Msg = err.Error()
			return resp, nil
		}
		if modelI == nil {
			mate.Code = ApiCode.REQ_DATA_ERROR
			mate.Msg = "数据不存在"
			return resp, nil
		}
		model = modelI.(models.` + modelName + `)
		if model.Id <= 0 {
			mate.Code = ApiCode.REQ_DATA_ERROR
			mate.Msg = "数据不存在"
			return resp, nil
		}
	} else {
		model = models.` + modelName + `{
			Id: modelbase.GenDbId(),
		}
	}
	model.ChannelId = in.GetChannelId()
	err := model.Save(l.ctx)
	if err != nil {
		mate.Code = ApiCode.DB_ERROR
		mate.Msg = err.Error()
		return resp, nil
	}
	returnData := modelClient.ReturnData(l.ctx, model, nil)
	if returnData != nil {
		resp.Data = returnData.(*` + pkgName + `.` + resDataName + `)
	}
	return resp, nil
`
	return str
}
