package http

import (
	"github.com/yagikota/gRPC_with_go/pkg/adapter/proto"
	"github.com/yagikota/gRPC_with_go/pkg/domain/service"
	"github.com/yagikota/gRPC_with_go/pkg/infra"
	"github.com/yagikota/gRPC_with_go/pkg/infra/mysql"
	"github.com/yagikota/gRPC_with_go/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitServer() *grpc.Server {
	mySQLConn := infra.NewMySQLConnector()
	studentRepository := mysql.NewStudentRepository(mySQLConn.Conn)
	studentService := service.NewStudentService(studentRepository)
	studentUsecase := usecase.NewUserUsecase(studentService)

	gs := grpc.NewServer()
	studentServer := NewStudentServer(gs, studentUsecase)
	proto.RegisterStudentServiceServer(gs, studentServer)
	healthCheckSever := NewHealthCheckSerer()
	proto.RegisterHealthCheckServiceServer(gs, healthCheckSever)
	reflection.Register(gs)
	return gs
}
