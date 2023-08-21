CTL = goctl
LINUXENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GOCMD = go
GOBUILD = $(GOCMD) build
PPROFTOOL = $(GOCMD) tool pprof

# 根据proto文件生成对应的数据
protoc:
	$(CTL) rpc protoc testdemo.proto --go_out=. --go-grpc_out=. --zrpc_out=.