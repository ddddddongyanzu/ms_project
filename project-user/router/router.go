package router

import (
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"test.com/project-common/discovery"
	"test.com/project-common/logs"
	"test.com/project-grpc/user/login"
	"test.com/project-user/config"
	"test.com/project-user/internal/interceptor"
	loginServiceV1 "test.com/project-user/pkg/service/login.service.v1"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	//rg := New()
	//rg.Route(&user.RouterUser{}, r)
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			login.RegisterLoginServiceServer(g, loginServiceV1.New())
		}}
	s := grpc.NewServer(
		// 1. OpenTelemetry 通过 StatsHandler 来进行 Telemetry 统计和追踪
		grpc.StatsHandler(
			otelgrpc.NewServerHandler(
				otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
			),
		),
		// 2. 自定义缓存拦截器走 UnaryInterceptor
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptor.New().CacheInterceptor(),
				// 如果需要，还可以链更多的拦截器
			),
		),
	)
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		log.Printf("grpc server started as: %s \n", c.Addr)
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight,
	}
	r := discovery.NewRegister(config.C.EtcdConfig.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
