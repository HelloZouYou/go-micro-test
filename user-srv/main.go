package main

import (
	"fmt"

	"github.com/HelloZouYou/go-micro-test/user-srv/basic"
	"github.com/HelloZouYou/go-micro-test/user-srv/basic/config"
	"github.com/HelloZouYou/go-micro-test/user-srv/handler"
	"github.com/HelloZouYou/go-micro-test/user-srv/model"
	s "github.com/HelloZouYou/go-micro-test/user-srv/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	etcd "github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	regitry := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(regitry),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	fmt.Println(etcdCfg)
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
