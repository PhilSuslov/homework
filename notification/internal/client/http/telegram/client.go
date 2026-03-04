package telegram

import (
	"context"

	"github.com/go-telegram/bot"
)

type client struct {
	bot *bot.Bot
}

func NewClient(bot *bot.Bot) *client {
	return &client{bot: bot}
}

func (c *client) SendMessage(ctx context.Context, chatID int64, text string) error {
	// text = escapeMarkdownV2(text)
	_, err := c.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
		// ParseMode: "MarkdownV2",
	})

	if err != nil {
		return err
	}

	return nil
}

// func escapeMarkdownV2(s string) string {
// 	special := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
// 	for _, ch := range special {
// 		s = strings.ReplaceAll(s, ch, `\`+ch)
// 	}
// 	return s
// }
