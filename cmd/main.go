package main

import (
	"context"
	"coworking/internal/adapters/http"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func gracefulShutdown(fiberServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found in root, using environment variables")
	}

	// Verificar que PORT está presente y es válido
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("Environment variable PORT is required")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT value: %s", portStr)
	}

	server := http.NewServer()
	done := make(chan bool, 1)

	go func() {
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("http server error: %s", err)
		}
	}()

	go gracefulShutdown(server, done)

	<-done
	log.Println("Graceful shutdown complete.")
}
