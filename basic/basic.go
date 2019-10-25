package basic

import (
	"github.com/HelloZouYou/go-micro-test/basic/config"
)

var (
	pluginFuncs []func()
)

// Options Options
type Options struct {
	EnableDB    bool
	EnableRedis bool
	cfgOps      []config.Option
}

// Option Option
type Option func(o *Option)

// Init Init
func Init(opts ...config.Option) {
	// 初始化配置
	config.Init(opts...)
	// 加载依赖配置的插件
	for _, f := range pluginFuncs {
		f()
	}
}

// Register Register
func Register(f func()) {
	pluginFuncs = append(pluginFuncs, f)
}
