package handlers

import (
	"context"
	"log"
	"go_antigcast/config"
	"go_antigcast/filters"
	"go_antigcast/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	text := update.Message.Text

	// Ambil config grup dari MongoDB
	cfg, err := config.GetGroupConfig(chatID)
	if err != nil {
		log.Println("Gagal ambil config:", err)
		return
	}

	// Jika anti-gcast off, skip
	if !cfg.Enabled {
		return
	}

	// Cek suspicious
	if filters.IsSuspicious(text, cfg.Blacklist, cfg.Whitelist) {
		log.Printf("Pesan mencurigakan di chat %d, menghapus pesan %d\n", chatID, update.Message.MessageID)

		// Hapus pesan
		err := utils.DeleteMessage(bot, chatID, update.Message.MessageID)
		if err != nil {
			log.Println("Gagal hapus pesan:", err)
		}
	}
}
