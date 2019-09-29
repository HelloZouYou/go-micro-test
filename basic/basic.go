package basic

import (
	"github.com/HelloZouYou/go-micro-test/basic/config"
	"github.com/HelloZouYou/go-micro-test/basic/db"
	"github.com/HelloZouYou/go-micro-test/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}