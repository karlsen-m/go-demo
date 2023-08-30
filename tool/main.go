package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"tool/cmd"
	"tool/gotoprotobuf"
	"tool/logic"
	"tool/models"
	"tool/protobuftomd"
	"tool/rpcvalidate"
)

const (
	ApiMd              = "apimd"
	ModelBuf           = "modelbuf"
	CreateModel        = "add_model"
	CreateLogic        = "add_logic"
	CreateLogicByModel = "model_logic"
	CreateVerify       = "add_verify"
	Protoc             = "protoc"
)

func main() {

	protoFileName, err := getProtoFile()
	if err != nil {
		fmt.Println("Error:  get proto file msg：" + err.Error())
		return
	}
	if protoFileName == "" {
		fmt.Println("Error: proto file not found")
		return
	}
	comment, err := getProtoComment(protoFileName)
	if err != nil {
		fmt.Println("Error:  get proto file comment msg：" + err.Error())
		return
	}
	pakegeName, err := getPakegeName(comment)
	if err != nil {
		fmt.Println("Error:  get proto file pakege err：" + err.Error())
		return
	}
	if pakegeName == "" {
		fmt.Println("Error: proto file package not found")
		return
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" {
			fmt.Println("-apimd: <serviceName> protobuf to markdown example apimd  Ping")
			fmt.Println("-modelbuf: <fileName> model struct to protobuf example modelbuf ./test.go")
			fmt.Println("-add_model: <modelName>  <fileName> modelName create model struct to file example add_model admin ./models/admin.go")
			fmt.Println("-add_logic:  <modelName> <logicType> <apiName> create logic struct to file example add_logic admin edit EditAdmin")
			fmt.Println("-model_logic:  <modelName> create logic struct to file example model_logic admin")
			fmt.Println("-add_verify:  create validate file by " + pakegeName + ".proto")
			fmt.Println("-protoc:  crrpc protoc " + pakegeName + ".proto --go_out=. --go-grpc_out=. --zrpc_out=.")
			return
		}
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <type> <option>")
		return
	}
	switch os.Args[1] {
	case Protoc:
		cmd.ZeroProtoc(pakegeName)
		break
	case CreateVerify:
		rpcvalidate.AddValidate(pakegeName, comment)
		break
	case ApiMd:
		if len(os.Args) < 3 {
			protobuftomd.ProtoBufToMd(comment, "")
		} else {
			protobuftomd.ProtoBufToMd(comment, os.Args[2])
		}
		break
	case ModelBuf:
		gotoprotobuf.Gotoprotobuf(os.Args[2])
		break
	case CreateModel:
		if len(os.Args) == 3 {
			models.CreateModel(os.Args[2], "")
		} else {
			models.CreateModel(os.Args[2], os.Args[3])
		}
		break
	case CreateLogic:
		if len(os.Args) >= 5 {
			messageDataMap := logic.GetMapMessageToDataWithCommen(comment)
			logic.CreateLogic(pakegeName, os.Args[2], os.Args[3], os.Args[4], messageDataMap)
		} else {
			fmt.Println("Usage: go run main.go <type> <option>")
		}
		break
	case CreateLogicByModel:
		if len(os.Args) >= 3 {
			logic.CreateLogicByModel(pakegeName, comment, os.Args[2])
		} else {
			fmt.Println("Usage: go run main.go <type> <option>")
		}

	default:
		fmt.Println("Usage: go run main.go <type> error")
		return
	}
}

func getProtoFile() (protoFileName string, err error) {
	dir, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		return
	}
	var protoFiles []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if err != nil {
		return
	}
	if len(protoFiles) <= 0 {
		return
	}
	protoFileName = protoFiles[0]
	return

}

func getPakegeName(comment []string) (string, error) {
	packageComment := ""
	for _, v := range comment {
		if strings.Contains(v, "package") {
			packageComment = v
			break
		}
	}
	if packageComment == "" {
		return "", errors.New("proto file package not found")
	}
	re1 := regexp.MustCompile(`package\s+`)
	replaced := re1.ReplaceAllString(packageComment, "")
	re2 := regexp.MustCompile(`;`)
	packageName := re2.ReplaceAllString(replaced, "")
	return packageName, nil
}

func getProtoComment(protoFilename string) ([]string, error) {
	// 读取 proto 文件内容
	protoContent, err := ioutil.ReadFile(protoFilename)
	if err != nil {

		return []string{}, errors.New(fmt.Sprintf("Error parsing file: %v\n", err))
	}
	protoComment := string(protoContent)

	// 匹配并替换单行注释
	//pattern1 := regexp.MustCompile(`//.*`)
	pattern1 := regexp.MustCompile(`\n\s*\/\/.*?\n`)
	protoComment = pattern1.ReplaceAllString(protoComment, "\n")

	// 匹配并替换段落注释
	pattern2 := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	protoComment = pattern2.ReplaceAllString(protoComment, "")

	// 匹配空行
	pattern3 := regexp.MustCompile(`\n{2,}`)
	protoComment = pattern3.ReplaceAllString(protoComment, "\n")

	comment := []string{}
	comment = strings.Split(protoComment, "\n")
	return comment, nil
}
