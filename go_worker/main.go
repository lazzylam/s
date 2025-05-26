package main

import (
	"log"
	"os"
	"go_antigcast/config"
	"go_antigcast/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := "anti_gcast"
	collName := "configs"
	botToken := os.Getenv("BOT_TOKEN")

	// Inisialisasi MongoDB
	config.InitMongo(mongoURI, dbName, collName)

	// Inisialisasi Telegram Bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Mulai polling update
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go handlers.HandleUpdate(bot, update)
	}
}
