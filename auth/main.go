package main

import (
	"fmt"

	"github.com/HelloZouYou/go-micro-test/auth/handler"
	"github.com/HelloZouYou/go-micro-test/auth/model"
	s "github.com/HelloZouYou/go-micro-test/auth/proto/auth"
	"github.com/HelloZouYou/go-micro-test/basic"
	"github.com/HelloZouYou/go-micro-test/basic/common"
	"github.com/HelloZouYou/go-micro-test/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/etcd"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var (
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用etcdv3注册
	micReg := etcdv3.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func initCfg() {
	source := etcd.NewSource(
		etcd.WithAddress("etcd:2379"),
		etcd.WithPrefix("/micro/auth_srv"),
		etcd.StripPrefix(true),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg]配置，cfg：%v", cfg)

	return
}
