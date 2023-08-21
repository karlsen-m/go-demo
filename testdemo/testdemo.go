package main

import (
  _ "utils/config"

	"flag"
	"fmt"
	"testdemo/internal/config"
	"testdemo/internal/server"
	"testdemo/internal/svc"
	"testdemo/testdemo"
	"utils/tracer"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	_ "utils/kafkabase"
	"go-common/tools/logger"
	"github.com/spf13/cast"
	"go-common/tools/helpers"
)

func main() {
	flag.Parse()

	var c config.Config
	c.ListenOn = helpers.GetConfigToString("ListenOn", "")
	c.Timeout = cast.ToInt64(helpers.GetConfigToString("Timeout", ""))
	c.CpuThreshold = 900
	c.ServiceConf.Name = helpers.GetConfigToString("Name", "")
	ctx := svc.NewServiceContext(c)
	srv := server.NewTestdemoServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		testdemo.RegisterTestdemoServer(grpcServer, srv)
	})
	defer s.Stop()
	s.AddUnaryInterceptors(tracer.TraceInterceptor())
	fmt.Printf("Starting rpc server at %s...", c.ListenOn)
	logger.Notice("微服务启动", "200", "启动成功", logger.H{})
	s.Start()
}

