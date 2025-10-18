package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/trianglehasfoursides/mathrock/auth"
)

func main() {
	server := echo.New()
	server.GET("/", index)
	server.GET("/status", status)

	api := server.Group("/api/v1", auth.Auth)
	api.GET("", nil)

	// for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
	}()

	server.Start("")
}
