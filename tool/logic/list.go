package logic

func createListLogic(pkgName, apiName, modelName, resDataName string) string {
	str := `
	mate := &` + pkgName + `.MetaRes{
		RequestId: tracing.GetRequestId(l.ctx),
		Code:      ApiCode.SUCCESS,
		Msg:       response.NewMessageConstant().SuccessMsg,
	}
	resp := &` + pkgName + `.` + apiName + `Res{
		Meta: mate,
	}
	resDataReturn := &` + pkgName + `.` + apiName + `Data{}
	page := cast.ToInt(in.GetPage())
	pageSize := cast.ToInt(in.GetPageSize())
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	search := make(map[string]interface{})
	modelClient := models.NewModelClient("` + modelName + `")
	total, err := modelClient.GetTotal(l.ctx, search)
	if err != nil {
		mate.Code = ApiCode.DB_ERROR
		mate.Msg = err.Error()
		return resp, nil
	}
	pageInfo := &` + pkgName + `.PageInfo{
		PageSize:  cast.ToString(pageSize),
		Page:      cast.ToString(page),
		Total:     cast.ToString(total),
		TotalPage: cast.ToString(math.Ceil(float64(total) / float64(pageSize))),
	}
	resDataReturn.PageInfo = pageInfo
	if total <= 0 {
		resp.Data = resDataReturn
		return resp, nil
	}
	modelList, err := modelClient.GetList(l.ctx, search, page, pageSize)
	if err != nil {
		mate.Code = ApiCode.DB_ERROR
		mate.Msg = err.Error()
		return resp, nil
	}
	if len(modelList) <= 0 {
		resp.Data = resDataReturn
		return resp, nil
	}
	for _, v := range modelList {
		returnData := modelClient.ReturnData(l.ctx, v, nil)
		if returnData != nil {
			resDataReturn.List = append(resDataReturn.List, returnData.(*` + pkgName + `.` + resDataName + `))
		}
	}
	resp.Data = resDataReturn
	return resp, nil
`
	return str
}
