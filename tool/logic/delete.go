package logic

func createDeleteLogic(pkgName, apiName, modelName string) string {
	modelNameBak := lowerFirstLetter(modelName)
	if modelNameBak == pkgName {
		modelNameBak += "M"
	}
	str := `
	meta := &` + pkgName + `.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &` + pkgName + `.` + apiName + `Res{
		Meta: meta,
	}
	id := cast.ToUint64(in.GetId())
	if id <= 0 {
		meta.Code = ApiCode.REQ_DATA_ERROR
		meta.Msg = "id不能为空"
		return resp, nil
	}
	` + modelNameBak + ` := models.` + modelName + `{}
	util, err := ` + modelNameBak + `.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	err = util.FirstById(l.ctx, &` + modelNameBak + `, cast.ToUint64(in.GetId()))
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	if !` + modelNameBak + `.IsExist() {
		meta.Code = ApiCode.REQ_DATA_ERROR
		meta.Msg = "数据不存在"
		return resp, nil
	}
	` + modelNameBak + `.DeletedAt = helpers.GetNowTime()
	err = util.SaveById(l.ctx, &` + modelNameBak + `, ` + modelNameBak + `.Id)
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	return resp, nil
`
	return str
}
