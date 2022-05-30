package main

import (
	"flag"
	"fmt"

	"github.com/ByteDance-camp/TickTalk/api/config"
	"github.com/ByteDance-camp/TickTalk/api/user"
	"github.com/ByteDance-camp/TickTalk/api/videomgr"
	"github.com/ByteDance-camp/TickTalk/api/videomgr/storage"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	user.InitApi(c, server)
	videomgr.InitApi(c, server)

	localStore, stop := storage.NewLocalHandler(c.LocalStore.Addr, c.LocalStore.Path)
	storage.RegisterHandler(localStore)
	go func() {
		<-stop
		logx.Infof("file server stopped")
		server.Stop()
		logx.Infof("api server stopped")
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
