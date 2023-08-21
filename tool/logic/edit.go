package logic

func createEditLogic(pkgName, apiName, modelName string) string {
	modelNameBak := lowerFirstLetter(modelName)

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
	` + modelNameBak + ` := models.` + modelName + `{}
	util, err := ` + modelNameBak + `.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	if id > 0 {
		err = util.FirstById(l.ctx, &` + modelNameBak + `, id)
		if err != nil {
			meta.Code = ApiCode.DB_ERROR
			meta.Msg = err.Error()
			return resp, nil
		}
		if !` + modelNameBak + `.IsExist() {
			meta.Code = ApiCode.REQ_DATA_ERROR
			meta.Msg = "不存在该用户"
			return resp, nil
		}
	} else {
		` + modelNameBak + `.Id = models.GenerateId()
	}
	` + modelNameBak + `.ChannelId = in.GetChannelId()
	err = util.SaveById(l.ctx, &` + modelNameBak + `, ` + modelNameBak + `.Id)
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	resp.Data = ` + modelNameBak + `.ReturnData(l.ctx)
	return resp, nil
`
	return str
}
