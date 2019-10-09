package main

import (
	"fmt"

	"github.com/HelloZouYou/go-micro-test/basic"
	"github.com/HelloZouYou/go-micro-test/basic/common"
	"github.com/HelloZouYou/go-micro-test/basic/config"
	"github.com/HelloZouYou/go-micro-test/orders-srv/handler"
	"github.com/HelloZouYou/go-micro-test/orders-srv/model"
	proto "github.com/HelloZouYou/go-micro-test/orders-srv/proto/orders"
	"github.com/HelloZouYou/go-micro-test/orders-srv/subscriber"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	etcd "github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.orders"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			// 初始化sub
			subscriber.Init()
		}),
	)

	// 侦听订单支付消息
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务
	err = proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}