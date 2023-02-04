package router

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/ms_project/project-user/config"
	loginServiceV1 "github.com/prynnekey/ms_project/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
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

// Init 初始化路由
func Init(r *gin.Engine) {
	// register := New()
	// register.Route(&user.RouterUser{}, r)
	for _, ro := range routers {
		ro.Route(r)
	}
}

// Register 注册路由
func Register(ro ...Router) {
	routers = append(routers, ro...)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

// 注册gRPC
func RegisterGrpc() *grpc.Server {
	g := gRPCConfig{
		Addr: config.AppConfig.GC.Addr,
		RegisterFunc: func(s *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(s, loginServiceV1.New())
		},
	}
	grpc := grpc.NewServer()
	g.RegisterFunc(grpc)

	// 启动监听
	listener, err := net.Listen("tcp", g.Addr)
	if err != nil {
		log.Println("gRPC listen error: ", err)
	}
	go func() {
		// 开启监听
		err := grpc.Serve(listener)
		if err != nil {
			log.Println("gRPC serve error: ", err)
			return
		}
	}()

	return grpc
}
