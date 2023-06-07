package logic

func createDeleteLogic(pkgName, apiName, modelName, resDataName string) string {
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
	model := modelI.(models.` + modelName + `)
	if model.Id <= 0 {
		mate.Code = ApiCode.REQ_DATA_ERROR
		mate.Msg = "数据不存在"
		return resp, nil
	}
	err = model.Delete(l.ctx)
	if err != nil {
		mate.Code = ApiCode.DB_ERROR
		mate.Msg = err.Error()
		return resp, nil
	}
	return resp, nil
`
	return str
}
