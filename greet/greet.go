package main

import (
	"flag"
	"fmt"
	"learnZero/greet/internal/config"
	"learnZero/greet/internal/handler"
	"learnZero/greet/internal/httpx"
	"learnZero/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 使用自己实现的router handler
	runOption := rest.WithRouter(httpx.NewRouter())
	ctx := svc.NewServiceContext(c)
	// 注入自己实现的handler
	server := rest.MustNewServer(c.RestConf, runOption)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
