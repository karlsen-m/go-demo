package rpcvalidate

import (
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"strings"
	"tool/logic"
	"tool/protobuftomd"
)

var (
	//notReqFields map[string]string
	dataType = []string{
		"float64",
		"float32",
		"int32",
		"int64",
		"uint32",
		"uint64",
		"int32",
		"int64",
		"uint32",
		"uint64",
		"int32",
		"int64",
		"bool",
		"string",
	}
)

func AddValidate(pkgName string, comment []string) {
	notReqFields := make(map[string]string)
	messageDataMap := logic.GetMapMessageToDataWithCommen(comment)
	rpcReqNameData := make([]string, 0)
	for _, v := range comment {
		if strings.Contains(v, "rpc") {
			serviceData := protobuftomd.ServiceStrSplit(v)
			if len(serviceData) > 0 {
				if serviceData["serviceName"] != "Ping" {
					rpcReqNameData = append(rpcReqNameData, serviceData["reqName"])
				}
			}
		}
		if strings.Contains(v, "message") {
			break
		}
	}
	str := `
package ` + pkgName + `

import (
	"errors"
)

var err error

`
	if len(rpcReqNameData) > 0 {
		for _, v := range rpcReqNameData {
			if reqData, ok := messageDataMap[v]; ok {
				reqValidateStr := `
func (in *` + v + `)Validate() error {
`
				for _, reqField := range reqData {
					fields := protobuftomd.FieldSplit(reqField)
					if len(fields) > 0 {
						if cast.ToString(fields[0]) == "string" {
							if strings.Contains(fields[3], "required=true") {
								reqValidateStr += `
	if in.Get` + strings.Title(fields[1]) + `() == "" {
		return errors.New("` + fields[1] + ` 不能为空")
	}`
							}
						} else if cast.ToString(fields[0]) == "repeated" {
							if strings.Contains(fields[4], "required=true") {
								reqValidateStr += `
	if len(in.Get` + strings.Title(fields[2]) + `()) <= 0 {
		return errors.New("` + fields[2] + ` 不能为空")
	}`
							} else if fields[1] != "string" {
								notReqFields[fields[1]] = ""
								reqValidateStr += `
	if len(in.Get` + strings.Title(fields[2]) + `()) > 0 {
		for _, v := range in.Get` + strings.Title(fields[2]) + `() {
			err = v.Validate()
			if err != nil {
				return err
			}
		}
	}`
							}
						} else if !protobuftomd.InArrayWithString(cast.ToString(fields[0]), dataType) {
							notReqFields[fields[0]] = ""
							reqValidateStr += `
	if err = in.Get` + strings.Title(fields[1]) + `().Validate(); err != nil {
		return err
	}`
						}
					}
				}
				reqValidateStr += `
	return nil
}
`
				str += reqValidateStr
			}
		}
	}
	if len(notReqFields) > 0 {
		notReqFieldsKeys := []string{}
		for k, _ := range notReqFields {
			notReqFieldsKeys = append(notReqFieldsKeys, k)
		}
		notReqStr := NotReqFieldsValidate(notReqFieldsKeys, messageDataMap)
		if notReqStr != "" {
			str = str + notReqStr
		}
	}
	err := ioutil.WriteFile("./"+pkgName+"/validate.go", []byte(str), os.ModePerm)
	if err != nil {
		fmt.Println("create validate file error：" + err.Error())
		return
	}
	fmt.Println("create validate file success")
	return
}

func NotReqFieldsValidate(rpcReqNameData []string, messageDataMap map[string][]string) string {
	notReqFields := make(map[string]string)
	str := ""
	if len(rpcReqNameData) > 0 {
		for _, v := range rpcReqNameData {
			if reqData, ok := messageDataMap[v]; ok {
				reqValidateStr := `
func (in *` + v + `)Validate() error {
`
				for _, reqField := range reqData {
					fields := protobuftomd.FieldSplit(reqField)
					if len(fields) > 0 {
						if cast.ToString(fields[0]) == "string" {
							if strings.Contains(fields[3], "required=true") {
								reqValidateStr += `
	if in.Get` + strings.Title(fields[1]) + `() == "" {
		return errors.New("` + fields[1] + ` 不能为空")
	}`
							}
						} else if cast.ToString(fields[0]) == "repeated" {
							if strings.Contains(fields[4], "required=true") {
								reqValidateStr += `
	if len(in.Get` + strings.Title(fields[2]) + `()) <= 0 {
		return errors.New("` + fields[2] + ` 不能为空")
	}`
							} else if fields[1] != "string" {
								//判断是否无限套娃，否就继续写
								if fields[1] != v {
									notReqFields[fields[1]] = ""
								}
								reqValidateStr += `
	if len(in.Get` + strings.Title(fields[2]) + `()) > 0 {
		for _, v := range in.Get` + strings.Title(fields[2]) + `() {
			err = v.Validate()
			if err != nil {
				return err
			}
		}
	}`
							}
						} else if !protobuftomd.InArrayWithString(cast.ToString(fields[0]), dataType) {
							//判断是否无限套娃，否就继续写
							if fields[0] != v {
								notReqFields[fields[0]] = ""
							}
							reqValidateStr += `
	if err = in.Get` + strings.Title(fields[1]) + `().Validate(); err != nil {
		return err
	}`
						}
					}
				}
				reqValidateStr += `
	return nil
}
`
				str += reqValidateStr
			}
		}
	}
	fmt.Println("NotReqFieldsValidate:", str)
	if len(notReqFields) > 0 {
		notReqFieldsKeys := []string{}
		for k, _ := range notReqFields {
			notReqFieldsKeys = append(notReqFieldsKeys, k)
		}
		notReqStr := NotReqFieldsValidate(notReqFieldsKeys, messageDataMap)
		if notReqStr != "" {
			str = str + notReqStr
		}
	}
	return str
}
