package basic

import (
	"github.com/HelloZouYou/go-micro-test/user-srv/basic/config"
	"github.com/HelloZouYou/go-micro-test/user-srv/basic/db"
)

// Init 基础组件初始化
func Init() {
	config.Init()
	db.Init()
}
