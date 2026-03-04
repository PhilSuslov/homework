package telegram

import (
	"bytes"
	"context"
	"embed"
	"html/template"

	"github.com/PhilSuslov/homework/notification/internal/client/http"
	"github.com/PhilSuslov/homework/notification/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

const chatID = 914563960 // поменять на настоящий

//go:embed templates/assembled_notification.tmpl
var templateShipFS embed.FS

//go:embed templates/paid_notification.tmpl
var templatePaidFS embed.FS

// type shipAssembledData model.ShipAssembled
// type orderPaidData model.OrderPaidEvent

var paidTemplate = template.Must(template.ParseFS(templatePaidFS, "templates/paid_notification.tmpl"))
var assemblyTemplate = template.Must(template.ParseFS(templateShipFS, "templates/assembled_notification.tmpl"))

type service struct {
	telegramClient http.TelegramClient
}

func NewService(telegramClient http.TelegramClient) *service {
	return &service{telegramClient: telegramClient}
}

func (s *service) SendAssembledNotification(ctx context.Context, uuid string, ship model.ShipAssembled) error {
	message, err := s.buildShipMessage(uuid, ship)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram Ship Notification message sent to chat", zap.Int("chat_id", chatID),
		zap.String("messgae", message))

	return nil
}

func (s *service) buildShipMessage(uuid string, ship model.ShipAssembled) (string, error) {
	data := model.ShipAssembled{
		Event_uuid:     ship.Event_uuid,
		Order_uuid:     ship.Order_uuid,
		User_uuid:      uuid,
		Build_time_sec: ship.Build_time_sec,
	}

	var buf bytes.Buffer
	err := assemblyTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *service) SendPaidNotification(ctx context.Context, uuid string, paid model.OrderPaidEvent) error {
	message, err := s.buildPaidMessage(uuid, paid)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram Order Paid Notification message sent to chat", zap.Int("chat_id", chatID),
		zap.String("messgae", message))

	return nil
}

func (s *service) buildPaidMessage(uuid string, paid model.OrderPaidEvent) (string, error) {
	data := model.OrderPaidEvent{
		Event_uuid:       paid.Event_uuid,
		Order_uuid:       paid.Order_uuid,
		User_uuid:        uuid,
		Payment_method:   paid.Payment_method,
		Transaction_uuid: paid.Transaction_uuid,
	}

	var buf bytes.Buffer
	err := paidTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
