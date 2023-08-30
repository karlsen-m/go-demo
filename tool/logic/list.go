package logic

import (
	"strings"
	"tool/protobuftomd"
)

func createListLogic(pkgName, apiName, modelName string, messageDataMap map[string][]string) string {
	searchData := ""
	modelNameBak := lowerFirstLetter(modelName)
	if modelNameBak == pkgName {
		modelNameBak += "M"
	}
	reqData, ok := messageDataMap[apiName+"Req"]
	if ok {
		for _, v := range reqData {
			fields := protobuftomd.FieldSplit(v)
			if len(fields) > 1 {
				if InArrayWithString(fields[0], dataType) {
					if fields[1] != "page" && fields[1] != "pageSize" {
						if searchData == "" {
							searchData += "search" + `["` + fields[1] + `"] = in.Get` + strings.Title(fields[1]) + "()\n"
						} else {
							searchData += "\tsearch" + `["` + fields[1] + `"] = in.Get` + strings.Title(fields[1]) + "()\n"
						}
					}
				} else if fields[0] == "repeated" {
					searchData += "\tsearch" + `["` + fields[2] + `"] = in.Get` + strings.Title(fields[2]) + "()\n"
				} else {
					searchData += "\tsearch" + `["` + fields[1] + `"] = in.Get` + fields[0] + "()\n"
				}
			}
		}
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
	err := in.Validate()
	if err != nil{
		meta.Code = ApiCode.REQ_DATA_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	data := &` + pkgName + `.` + apiName + `Data{}
	page := cast.ToInt(in.GetPage())
	pageSize := cast.ToInt(in.GetPageSize())
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	search := make(map[string]interface{})
	` + searchData + `
	` + modelNameBak + ` := models.` + modelName + `{}
	util, err := ` + modelNameBak + `.NewUtil()
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	total, err := util.GetTotal(l.ctx, &` + modelNameBak + `, search)
	if err != nil {
		meta.Code = ApiCode.DB_ERROR
		meta.Msg = err.Error()
		return resp, nil
	}
	data.PageInfo = &` + pkgName + `.PageInfo{
		Total:     cast.ToString(total),
		Page:      cast.ToString(page),
		PageSize:  cast.ToString(pageSize),
		TotalPage: util.GetTotalPage(total, pageSize),
	}
	resp.Data = data
	if total <= 0 {
		return resp, nil
	}
	list := make([]models.` + modelName + `, 0)
	err = util.GetList(l.ctx, &` + modelNameBak + `, search, &list, page, pageSize, "id desc")
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
`
	return str
}
