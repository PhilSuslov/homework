package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/PhilSuslov/homework/notification/internal/app"
	"github.com/PhilSuslov/homework/notification/internal/config"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

const configPath = "../../deploy/compose/notification/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось запустить приложение", zap.Error(err))
		return
	}

	logger.Info(appCtx, "✅ Сервер запущен")
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}
