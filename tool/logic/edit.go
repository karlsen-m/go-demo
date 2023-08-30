package logic

import (
	"strings"
	"tool/protobuftomd"
)

func createEditLogic(pkgName, apiName, modelName string, messageDataMap map[string][]string) string {
	modelData := ""
	modelNameBak := lowerFirstLetter(modelName)
	if modelNameBak == pkgName {
		modelNameBak += "M"
	}
	reqData, ok := messageDataMap[apiName+"Req"]
	if ok {
		for _, v := range reqData {
			fields := protobuftomd.FieldSplit(v)
			if len(fields) > 1 {
				if fields[1] != "id" {
					if InArrayWithString(fields[0], dataType) {
						if modelData == "" {
							modelData += modelNameBak + "." + strings.Title(fields[1]) + " = in.Get" + strings.Title(fields[1]) + "()\n"
						} else {
							modelData += "\t" + modelNameBak + "." + strings.Title(fields[1]) + " = in.Get" + strings.Title(fields[1]) + "()\n"
						}

					} else if fields[0] == "repeated" {
						modelData += "\t" + modelNameBak + "." + strings.Title(fields[2]) + " = in.Get" + strings.Title(fields[2]) + "()\n"
					} else {
						modelData += "\t" + modelNameBak + "." + strings.Title(fields[1]) + " = in.Get" + fields[0] + "()\n"
					}
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
` + modelData + `
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
