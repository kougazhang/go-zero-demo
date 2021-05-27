package main

import (
	"blog/api/internal/types"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"blog/api/internal/config"
	"blog/api/internal/handler"
	"blog/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (i int, i2 interface{}) {
		return http.StatusBadRequest, types.CommResp{
			Ok:    false,
			Error: err.Error(),
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
