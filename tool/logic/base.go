package logic

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	Edit   = "edit"
	Detail = "detail"
	Delete = "delete"
	List   = "list"
)

func CreateLogic(pkgName, modelName, resDataName, logicType, apiName string) {

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
		str = createEditLogic(pkgName, apiName, modelName, resDataName)
		importStr += `"utils/modelbase"`
		break
	case Detail:
		str = createDetailLogic(pkgName, apiName, modelName, resDataName)
		break
	case Delete:
		str = createDeleteLogic(pkgName, apiName, modelName, resDataName)
		break
	case List:
		str = createListLogic(pkgName, apiName, modelName, resDataName)
		importStr += `"math"`

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
