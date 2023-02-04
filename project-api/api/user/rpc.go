package user

import (
	"log"

	login_service_v1 "github.com/prynnekey/ms_project/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var LoginServiceClient login_service_v1.LoginServiceClient

// 初始化grpc链接
func InitRpcUserClient() {
	conn, err := grpc.Dial(":8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("grpc.Dial err:", err)
	}

	LoginServiceClient = login_service_v1.NewLoginServiceClient(conn)
}
