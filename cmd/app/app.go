package app

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"websirk/config"
	"websirk/internal/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run(cfg *config.Config) error {
	//Server
	server := server.NewServer(cfg)
	err := server.Start()
	if err != nil {
		slog.Error("can't start a server %w", err)
		return err
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		slog.Error("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		slog.Error("app - Run - httpServer.Shoutdown: %w", err)
	}
	return nil
}
