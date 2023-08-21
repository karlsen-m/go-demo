package logic

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

const (
	Edit   = "edit"
	Detail = "detail"
	Delete = "delete"
	List   = "list"
)

func CreateLogicByModel(pkgName, modelName string) {
	//rpc EditTag(EditTagReq) returns(EditTagRes){};
	//    rpc GetTagDetail(GetTagReq) returns(GetTagRes){};
	//    rpc DelTag(DelTagReq) returns(DelTagRes){};
	//    rpc GetTagList(GetTagReq) returns(GetTagRes){};
	modelNameCapital := strings.Title(modelName)
	CreateLogic(pkgName, modelName, Edit, "Edit"+modelNameCapital)
	CreateLogic(pkgName, modelName, Detail, "Get"+modelNameCapital+"Detail")
	CreateLogic(pkgName, modelName, Delete, "Del"+modelNameCapital)
	CreateLogic(pkgName, modelName, List, "Get"+modelNameCapital+"List")
}

func CreateLogic(pkgName, modelName, logicType, apiName string) {

	importStr := `
	"context"
	"github.com/spf13/cast"
	ApiCode "go-common/code"
	"go-common/constant/response"
	"go-common/tools/tracing"
	"` + pkgName + `/models"
`

	str := ""
	switch logicType {
	case Edit:
		str = createEditLogic(pkgName, apiName, modelName)
		//importStr += `"utils/modelbase"`
		break
	case Detail:
		str = createDetailLogic(pkgName, apiName, modelName)
		break
	case Delete:
		str = createDeleteLogic(pkgName, apiName, modelName)
		break
	case List:
		str = createListLogic(pkgName, apiName, modelName)
		//importStr += `"math"`

		break
	default:
		fmt.Println("logicType error")
		return
	}
	lower := strings.ToLower(apiName)
	fileName := "./internal/logic/" + lower + "logic.go"
	protoContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}
	protoComment := string(protoContent)
	protoComment = strings.Replace(protoComment, `"context"`, importStr, -1)
	protoComment = strings.Replace(protoComment, "return &"+pkgName+"."+apiName+"Res{}, nil", "", -1)
	protoComment = strings.Replace(protoComment, "// todo: add your logic here and delete this line", str, -1)
	err = ioutil.WriteFile(fileName, []byte(protoComment), os.ModePerm)
	if err != nil {
		fmt.Println("create file error：" + err.Error())
		return
	}
	fmt.Println("create to logic success")

}

func lowerFirstLetter(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	if unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}

	return string(runes)
}
