package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	adaptorHTTP "github.com/yagikota/gRPC_with_go/pkg/adapter/http"
	"github.com/yagikota/gRPC_with_go/pkg/config"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := config.LoadConfig().HTTPInfo.Addr
	gs := adaptorHTTP.InitServer()
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// graceful shutdown
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	go func() {
		log.Printf("server is running! addr: %v", addr)
		gs.Serve(listener)
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	gs.GracefulStop()

	log.Println("Server exiting")
}
