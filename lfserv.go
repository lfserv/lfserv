package main

import (
	"github.com/apex/log"
	"github.com/gramework/gramework"
	"lfserv/api/store"
	"lfserv/storages/content/filesystem_store"
	"lfserv/storages/meta/boltdb_store"
	"lfserv/storages/meta/redis_store"
	"os"
	"strings"
)

const (
	envBindAddr     = "LFS_BIND"
	envMetaStore    = "LFS_META_STORE"
	envContentStore = "LFS_CONTENT_STORE"
)

var (
	bindAddr     string
	metaStore    store.MetaStore
	contentStore store.ContentStore
)

func init() {
	bindAddr = os.Getenv(envBindAddr)
	mStore := strings.ToLower(os.Getenv(envMetaStore))
	cStore := strings.ToLower(os.Getenv(envContentStore))
	if bindAddr == "" {
		bindAddr = ":8080"
	}
	switch mStore {
	case "redis":
		metaStore = &redis_store.RedisStore{}
	default:
		metaStore = &boltdb_store.BoltStore{}
	}
	switch cStore {
	default:
		contentStore = &filesystem_store.FileSystemStore{}
	}
}

func main() {
	err := metaStore.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = contentStore.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	app := gramework.New()
	app.GET("/", "hello")
	app.GET("/:name", func(ctx *gramework.Context) {
		name := ctx.RouteArg("name")
		ctx.WriteString("hello, " + name)
	})
	app.ListenAndServe(bindAddr)
}
