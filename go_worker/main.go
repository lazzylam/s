package main

import (
    "context"
    "fmt"
    "strings"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var blacklist = []string{"tmo", "vcs"}
var whitelist = []string{"tmobile"}

func main() {
    // Simulasi: connect ke Mongo
    ctx := context.TODO()
    client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    db := client.Database("anti_gcast")
    coll := db.Collection("config")

    // Simulasi: pesan masuk
    pesanMasuk := "Ayo vcs yuk @cewek"

    // Simulasi ambil config dari DB
    chatID := 123456
    config := getConfigFromMongo(ctx, coll, chatID)

    if isSuspicious(pesanMasuk, config.Blacklist, config.Whitelist) {
        fmt.Println("Pesan mencurigakan! Harus dihapus.")
        // call Telegram API untuk delete
    } else {
        fmt.Println("Pesan aman.")
    }
}
