package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"powergate.com/hr_timesheet/api/pb"
	"powergate.com/hr_timesheet/cmd/server/configs"
	"powergate.com/hr_timesheet/component"
	"powergate.com/hr_timesheet/internal/middleware"
	"powergate.com/hr_timesheet/internal/services"
)

func runGrpcServer(appContext component.AppContext) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(middleware.RecoverInterceptor))),
	)
	serverInstance := services.NewHrTimeSheet(appContext)

	pb.RegisterHrTimeSheetServiceServer(grpcServer, serverInstance)
	reflection.Register(grpcServer)
	conf := appContext.GetAppConfig()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCSERVERPORT))
	if err != nil {
		log.Fatal("cannot create listener")
	}
	log.Println("===== grpc start with port: ", conf.GRPCSERVERPORT, " =====")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func main() {
	config := configs.LoadConfig()

	appCtx := component.NewAppContext(config)

	runGrpcServer(appCtx)

	fmt.Println("hello1")

}
