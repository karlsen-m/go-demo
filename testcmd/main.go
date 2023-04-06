package main

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/color"
	"os"
)

//const (
//	TestServerTip     = `测试服`
//	TempProdServerTip = `临时生产服`
//	ProdServerTip     = `生产服`
//	QuitTip           = `退出`
//)
//
//func main() {
//	prompt := promptui.Select{
//		Label: "请选择发布服",
//		Items: []string{TestServerTip, ProdServerTip, QuitTip},
//	}
//
//	_, result, err := prompt.Run()
//	if err != nil {
//		fmt.Printf("Prompt failed %v\n", err)
//		return
//	}
//	t := time.Now().Format("20060102_150405")
//	file := fmt.Sprintf("build-%s.zip", t)
//	server := ""
//	serverPwd := ""
//	switch result {
//	case TestServerTip:
//		server = "root@112.74.160.112"
//		serverPwd = "/home/nginx/qymall/go"
//		break
//	//case TempProdServerTip:
//	//	server = "root@119.29.242.196"
//	//	serverPwd = "/www/wwwroot/efulihui.wxmas.com/go"
//	//	break
//	case ProdServerTip:
//		server = "root@8.134.95.124"
//		serverPwd = "/home/nginx/qymall/go"
//		break
//	case QuitTip:
//		os.Exit(1)
//	}
//	gitPull()
//	buildCode()
//	zipCode(file)
//	removeFile("./bin/run")
//	uploadCode(server, serverPwd, file)
//	runServerSh(server, serverPwd, file)
//	removeFile(file)
//	fmt.Println(color.Green("发布完毕,收工🫡!"))
//}
//
//func removeFile(file string) {
//	fmt.Println(color.Yellow("删除本地文件:\n" + file))
//	msg, _err := runCommand("./", "/bin/bash", "-c", fmt.Sprintf("rm -rf %s", file))
//	if _err != nil {
//		fmt.Println(color.Red("删除本地文件失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func runServerSh(server, serverPwd, file string) {
//	fmt.Println(color.Green("运行服务端脚本🫠🫠🫠"))
//	msg, _err := runCommand("./", "/bin/bash", "-c", fmt.Sprintf("ssh %s \"%s/deploy.sh %s\"", server, serverPwd, file))
//	if _err != nil {
//		fmt.Println(color.Red("运行服务端脚本失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func uploadCode(server, serverPwd, file string) {
//	fmt.Println(color.Green("上传代码🥰🥰🥰"))
//	msg, _err := runCommand("./", "/bin/bash", "-c", fmt.Sprintf("scp %s %s:%s/zips/", file, server, serverPwd))
//	if _err != nil {
//		fmt.Println(color.Red("上传代码失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func zipCode(file string) {
//	fmt.Println(color.Green("开始压缩程序🥶🥶🥶"))
//	msg, _err := runCommand("./", "/bin/bash", "-c", fmt.Sprintf("zip -q -r %s ./bin/run", file))
//	if _err != nil {
//		fmt.Println(color.Red("压缩程序失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func buildCode() {
//	fmt.Println(color.Green("开始编译代码🤩🤩🤩"))
//	msg, _err := runCommand("./", "/bin/bash", "-c", "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/run ./")
//	if _err != nil {
//		fmt.Println(color.Red("编译代码失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func gitPull() {
//	fmt.Println(color.Green("开始拉取代码🥳🥳🥳"))
//	msg, _err := runCommand("./", "/bin/bash", "-c", "git pull")
//	if _err != nil {
//		fmt.Println(color.Red("拉取代码失败;信息:\n" + _err.Error()))
//		os.Exit(1)
//	}
//	fmt.Println(color.White(msg))
//}
//
//func runCommand(path, name string, arg ...string) (msg string, err error) {
//	cmd := exec.Command(name, arg...)
//	var out bytes.Buffer
//	var stderr bytes.Buffer
//	cmd.Stdout = &out
//	cmd.Stderr = &stderr
//	cmd.Dir = path
//	err = cmd.Run()
//	if err != nil {
//		err = errors.New(fmt.Sprint(err) + ": " + stderr.String())
//	}
//	msg = out.String()
//	return
//}

func main() {
	var fileName string

	flag.StringVar(&fileName, "f", "", "proto文件名")
	flag.Parse()
	if fileName == "" {
		fmt.Println(color.Red("请输入proto文件名"))
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("fileName：%s", fileName))
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println(color.Red(err.Error()))
		os.Exit(-1)
	}
}
