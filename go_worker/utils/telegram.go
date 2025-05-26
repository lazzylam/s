package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) error {
	del := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(del)
	return err
}
