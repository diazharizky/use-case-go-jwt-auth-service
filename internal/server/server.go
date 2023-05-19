package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/diazharizky/use-case-go-jwt-auth-service/config"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing"
)

func init() {
	config.Global.SetDefault("server.host", "0.0.0.0")
	config.Global.SetDefault("server.port", "8080")
}

func Start() {
	host := config.Global.GetString("server.host")
	port := config.Global.GetString("server.port")

	addr := fmt.Sprintf("%s:%s", host, port)
	router := routing.NewRouter()

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		fmt.Println()
		log.Printf("Server is listening on %s", addr)

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	svrShutDownTimeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), svrShutDownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting!")
}
