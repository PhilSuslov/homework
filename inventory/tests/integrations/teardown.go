package integrations

import (
	"context"

	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func teardownTestEnvironment(ctx context.Context, env *TestEnvironment) {
	log := logger.Logger()
	log.Info(ctx, "🧹 Очистка тестового окружения...")

	cleanupTestEnvironment(ctx, env)

	log.Info(ctx, "✅ Тестовое окружение успешно очищено")
}

func cleanupTestEnvironment(ctx context.Context, env *TestEnvironment) {
	if env.App != nil {
		if err := env.App.Terminate(ctx); err != nil {
			logger.Error(ctx, "Не удалось остановить контейнер приложения", zap.Error(err))
		} else {
			logger.Info(ctx, "🔴 Контейнер приложения остановлен")
		}
	}

	if env.Mongo != nil {
		if err := env.Mongo.Terminate(ctx); err != nil {
			logger.Error(ctx, "Не удалось остановить контейнер MongoDB", zap.Error(err))
		} else {
			logger.Info(ctx, "🔴 Контейнер MongoDB остановлен")
		}
	}

	if env.Network != nil {
		if err := env.Network.Remove(ctx); err != nil {
			logger.Error(ctx, "Не удалось удалить сеть", zap.Error(err))
		} else {
			logger.Info(ctx, "🔴 Сеть удалена")
		}
	}
}
