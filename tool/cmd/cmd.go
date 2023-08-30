package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func ZeroProtoc(pakegeName string) {
	cmd := exec.Command("goctl", "rpc", "protoc", pakegeName+".proto", "--go_out=.", "--go-grpc_out=.", "--zrpc_out=.")

	// 设置命令的标准输出和错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("命令执行出错:", err)
		return
	}
}
