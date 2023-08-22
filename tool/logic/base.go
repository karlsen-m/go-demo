package logic

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"tool/protobuftomd"
	"unicode"
)

const (
	Edit   = "edit"
	Detail = "detail"
	Delete = "delete"
	List   = "list"
)

var dataType = []string{
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

func CreateLogicByModel(pkgName string, comment []string, modelName string) {
	messageDataMap := GetMapMessageToDataWithCommen(comment)
	modelNameCapital := strings.Title(modelName)
	CreateLogic(pkgName, modelName, Edit, "Edit"+modelNameCapital, messageDataMap)
	CreateLogic(pkgName, modelName, Detail, "Get"+modelNameCapital+"Detail", messageDataMap)
	CreateLogic(pkgName, modelName, Delete, "Del"+modelNameCapital, messageDataMap)
	CreateLogic(pkgName, modelName, List, "Get"+modelNameCapital+"List", messageDataMap)
}

func CreateLogic(pkgName, modelName, logicType, apiName string, messageDataMap map[string][]string) {

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
		str = createEditLogic(pkgName, apiName, modelName, messageDataMap)
		//importStr += `"utils/modelbase"`
		break
	case Detail:
		str = createDetailLogic(pkgName, apiName, modelName)
		break
	case Delete:
		str = createDeleteLogic(pkgName, apiName, modelName)
		break
	case List:
		str = createListLogic(pkgName, apiName, modelName, messageDataMap)
		//importStr += `"math"`

		break
	default:
		fmt.Println("logicType error")
		return
	}
	if str == "" {
		fmt.Println("create logic fail logicComment not found")
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

func GetMapMessageToDataWithCommen(comment []string) (messageDataMap map[string][]string) {
	messageStartNum := 0
	for i, v := range comment {
		if strings.Contains(v, "message") {
			messageStartNum = i
			break
		}
	}
	messageComment := make([]string, len(comment[messageStartNum:]))
	_ = copy(messageComment, comment[messageStartNum:])

	messageData := [][]string{}
START:
	if len(messageComment) > 0 {
		messageNum := 0
		for i, v := range messageComment {
			if strings.Contains(v, "message") {
				messageNum++
			}
			if strings.Contains(v, "}") {
				messageNum--
				if messageNum == 0 {
					//一个大的message结束
					extractData := make([]string, len(messageComment[:i+1]))
					_ = copy(extractData, messageComment[:i+1])
					data := protobuftomd.ExtractMessagesData("", extractData)
					messageData = append(messageData, data...)
					messageComment = messageComment[i+1:]
					goto START
				}
			}
		}
	}
	messageDataMap = make(map[string][]string)
	for _, v := range messageData {
		if len(v) > 4 {
			fields := make([]string, len(v[3:])-1)
			_ = copy(fields, v[3:])
			messageDataMap[v[1]] = fields
		}
	}
	return
}

// 判断值是否存在于数组里(string版)
func InArrayWithString(val string, arr []string) bool {
	for _, i := range arr {
		if i == val {
			return true
		}
	}
	return false
}
