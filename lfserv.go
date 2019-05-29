package main

import (
	"github.com/apex/log"
	"github.com/gramework/gramework"
	"lfserv/storage/filesystem_store"
	"lfserv/storage/redis_store"
	"lfserv/store"
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
		log.Fatalf("%s unknown or empty", envMetaStore)
	}
	switch cStore {
	case "filesystem":
		contentStore = &filesystem_store.FileSystemStore{}
	default:
		log.Fatalf("%s unknown or empty", envContentStore)
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
