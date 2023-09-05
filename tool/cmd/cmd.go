package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	Protoc = "protoc"
)

func Cmd(pakegeName string, cmdType string) {
	switch cmdType {
	case Protoc:
		BuildProtoc(pakegeName)
		break
	}
}

func BuildProtoc(pakegeName string) {
	frame := GetFrame(pakegeName)
	if frame == "" {
		return
	}
	if frame == "zero" {
		cmd := exec.Command("goctl", "rpc", "protoc", pakegeName+".proto", "--go_out=.", "--go-grpc_out=.", "--zrpc_out=.")
		// 设置命令的标准输出和错误输出
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// 执行命令
		err := cmd.Run()
		if err != nil {
			fmt.Println("goctl命令执行出错:", err)
			return
		}
	}
	cmd := exec.Command("truss", "./"+pakegeName+".proto")
	// 设置命令的标准输出和错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("truss 命令执行出错:", err)
		return
	}
	serviceGrpcPath := `./` + pakegeName + `-service/svc/client/grpc/imclient.go`
	serviceGrpcPoolPath := `./` + pakegeName + `-service/svc/client/grpc/poolimclient.go`
	serviceProtoFile := pakegeName + `.proto`
	servicePBFile := pakegeName + `.pb.go`

	commonPBPath := `../go-common/services/` + pakegeName + `/PB`
	commonGrpcPath := `../go-common/services/` + pakegeName + `/client/grpc`

	commonProtoFile := commonPBPath + `/` + pakegeName + `.proto`
	commonPBFile := commonPBPath + `/` + pakegeName + `.pb.go`

	commonGrpcFile := commonGrpcPath + `/imclient.go`
	commonGrpcPoolFile := commonGrpcPath + `/poolimclient.go`

	mkdirPB := exec.Command("mkdir", "-p", commonPBPath)
	_, err = mkdirPB.CombinedOutput()
	if err != nil {
		fmt.Println("mkdir err :", err)
		return
	}
	mkdirGrpc := exec.Command("mkdir", "-p", commonGrpcPath)
	_, err = mkdirGrpc.CombinedOutput()
	if err != nil {
		fmt.Println("mkdir err :", err)
		return
	}
	cpProto := exec.Command("cp", serviceProtoFile, commonProtoFile)
	_, err = cpProto.CombinedOutput()
	if err != nil {
		fmt.Println("cp .proto file err :", err)
		return
	}
	cpPB := exec.Command("cp", servicePBFile, commonPBFile)
	_, err = cpPB.CombinedOutput()
	if err != nil {
		fmt.Println("cp .pb.go file err :", err)
		return
	}
	cpGrpc := exec.Command("cp", serviceGrpcPath, commonGrpcFile)
	_, err = cpGrpc.CombinedOutput()
	if err != nil {
		fmt.Println("cp imclient.go file err :", err)
		return
	}
	cpGrpcPool := exec.Command("cp", serviceGrpcPoolPath, commonGrpcPoolFile)
	_, err = cpGrpcPool.CombinedOutput()
	if err != nil {
		fmt.Println("cp poolimclient.go file err :", err)
		return
	}
	if frame == "zero" {
		rmPB := exec.Command("rm", servicePBFile)
		_, err = rmPB.CombinedOutput()
		if err != nil {
			fmt.Println("rm .pb.go file err :", err)
			return
		}
		rmDir := exec.Command("rm", "-r", "./"+pakegeName+"-service")
		_, err = rmDir.CombinedOutput()
		if err != nil {
			fmt.Println("rm "+pakegeName+"-service dir err :", err)
			return
		}
	}
	fmt.Println("=====部署" + pakegeName + "到公共模块完成=====")
}

func GetFrame(pakegeName string) (frame string) {
	frame = ""
	currentDir, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}

	// 遍历当前目录中的文件
	for _, file := range currentDir {
		if file.Name() == pakegeName+".go" || "main.go" == file.Name() {
			frame = "zero"
			return
		}
	}
	frame = "gokit"
	return
}
