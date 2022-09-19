package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/yagikota/gRPC_with_go/pkg/adapter/proto"
	"google.golang.org/grpc"
)

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

type healthCheckSever struct {
	proto.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckSerer() *healthCheckSever {
	return &healthCheckSever{}
}

// // for health check
// //
// //	{
// //	    "message": "Hello, C Team. you've requested: /health_check"
// //	}
// //
// // will return
func (hs *healthCheckSever) HealthCheck(ctx context.Context, req *proto.HealthcheckRequest) (*proto.HealthcheckResponse, error) {
	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, errors.New("no method string for the server context")
	}
	fmt.Println(method)
	return &proto.HealthcheckResponse{
		Message: fmt.Sprintf("Hello! You've requested: %s", method),
	}, nil
}
