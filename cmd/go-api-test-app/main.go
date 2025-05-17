package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nikolaevnikita/go-api-test-app/internal/app"
	"github.com/nikolaevnikita/go-api-test-app/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(cancel)


	config := config.ReadConfig()
	app.NewApp(config).Start(ctx)
}

func gracefulShutdown(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigChan
	cancel()
}

// TODO:
// 1. Добавить логирование
// 2. Добавить авторизацию
// 3. Связать Task c User