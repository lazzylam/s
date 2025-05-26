package models

type GroupConfig struct {
	ChatID     int64    `bson:"chat_id"`
	Enabled    bool     `bson:"enabled"`
	Blacklist  []string `bson:"blacklist"`
	Whitelist  []string `bson:"whitelist"`
}
