package main

import (
	"fmt"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/options"
	"github.com/micro/go-micro/data/store"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/data/store/etcd"
)

var (
	// 根据需要改成可配置的app列表
	apps = []string{"micro"}
	record = 
)

// Service Service
type Service struct{}

func main() {
	// 灾难恢复
	defer func() {
		if r := recover(); r != nil {
			log.Logf("[main] Recovered in f %v", r)
		}
	}()

	// 加载并侦听配置文件
	err := loadAndWatchConfigFile()
	if err != nil {
		log.Fatal(err)
	}

}

func loadAndWatchConfigFile() (err error) {
	// 加载每个应用的配置文件
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath("./conf/" + app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadAndWatchConfigFile] 加载应用配置文件 异常，%s", err)
			return err
		}
	}
	for _, node := range config.Map() {
		fmt.Println(node)
	}
	store := etcd.NewStore(options.WithString("etcd:2379"))
	fmt.Println(store.Read("/haha"))
	store.Write(&store.Record{
		"heihei"

	})
	fmt.Println(store.Read("/haha"))

	// 侦听文件变动
	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadAndWatchConfigFile] 开始侦听应用配置文件变动 异常，%s", err)
		return err
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] 侦听应用配置文件变动 异常， %s", err)
				return
			}

			log.Logf("[loadAndWatchConfigFile] 文件变动，%s", string(v.Bytes()))
		}
	}()

	return
}
