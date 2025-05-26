import os
from pyrogram import Client, filters
from pymongo import MongoClient

API_ID = int(os.getenv("API_ID"))
API_HASH = os.getenv("API_HASH")
BOT_TOKEN = os.getenv("BOT_TOKEN")
MONGO_URI = os.getenv("MONGO_URI")

app = Client("admin_bot", api_id=API_ID, api_hash=API_HASH, bot_token=BOT_TOKEN)
mongo = MongoClient(MONGO_URI)
db = mongo["anti_gcast"]
coll = db["configs"]

def get_config(chat_id):
    cfg = coll.find_one({"chat_id": chat_id})
    if not cfg:
        cfg = {"chat_id": chat_id, "enabled": False, "blacklist": [], "whitelist": []}
        coll.insert_one(cfg)
    return cfg

def update_config(chat_id, update):
    coll.update_one({"chat_id": chat_id}, {"$set": update})

@app.on_message
