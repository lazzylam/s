package config

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_antigcast/models"
)

var (
	client     *mongo.Client
	ConfigColl *mongo.Collection
)

func InitMongo(uri string, dbName string, collName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	ConfigColl = client.Database(dbName).Collection(collName)
}

// Ambil konfigurasi grup
func GetGroupConfig(chatID int64) (*models.GroupConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var config models.GroupConfig
	err := ConfigColl.FindOne(ctx, map[string]interface{}{
		"chat_id": chatID,
	}).Decode(&config)

	if err != nil {
		// Jika belum ada, return default
		return &models.GroupConfig{
			ChatID:    chatID,
			Enabled:   false,
			Blacklist: []string{},
			Whitelist: []string{},
		}, nil
	}
	return &config, nil
}
