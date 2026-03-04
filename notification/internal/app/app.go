package app

import (
	"context"

	orderPaidConsumer "github.com/PhilSuslov/homework/notification/internal/service/consumer/order_paid_consumer"
	orderAssembledConsumer "github.com/PhilSuslov/homework/notification/internal/service/consumer/order_assembled_consumer"
	"github.com/PhilSuslov/homework/notification/internal/config"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

   // ▶️ Запускаем order.paid consumer
    go func() {
        svc := orderPaidConsumer.NewService(
            a.diContainer.OrderPaidConsumer(),
            a.diContainer.OrderPaidDecoder(),
			a.diContainer.TelegramPaidService(ctx),
        )

        if err := svc.RunConsumer(ctx); err != nil {
            errCh <- err
        }
    }()

    // ▶️ Запускаем order.assembled consumer
    go func() {
        svc := orderAssembledConsumer.NewService(
            a.diContainer.OrderAssembledConsumer(),
            a.diContainer.OrderAssembledDecoder(),
			a.diContainer.TelegramAssembledService(ctx),
        )

        if err := svc.RunConsumer(ctx); err != nil {
            errCh <- err
        }
    }()


	select {
	case <-ctx.Done():
		logger.Info(ctx, "Shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "Component crashed, shutting down", zap.Error(err))
		cancel()
		<-ctx.Done()
		return err
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initTelegramBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initTelegramBot(ctx context.Context) error {
	telegramBot := a.diContainer.TelegramBot(ctx)

	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			logger.Info(ctx, "chat id", zap.Int64("chat_id", update.Message.Chat.ID))

			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚚 Order Bot активирован! Теперь вы будете получать уведомления о новых заказах",
			})
			if err != nil {
				logger.Error(ctx, "Failed to send activation message", zap.Error(err))
			}
		})

	go func() {
		logger.Info(ctx, "🤖 Telegram bor started...")
		telegramBot.Start(ctx)
	}()

	return nil
}

