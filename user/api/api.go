package main

import (
	"flag"
	"fmt"

	"github.com/ByteDance-camp/TickTalk/user/api/internal/config"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/handler"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}