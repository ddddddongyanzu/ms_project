package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-project/config"
	"test.com/project-project/router"
)

func main() {
	r := gin.Default()
	// 路由
	router.InitRouter(r)
	// grpc 服务注册
	grpc := router.RegisterGrpc()
	// grpc 服务注册到 etcd
	router.RegisterEtcdServer()
	stop := func() {
		grpc.Stop()
	}
	// 初始化 rpc 调用
	router.InitUserRpc()
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
