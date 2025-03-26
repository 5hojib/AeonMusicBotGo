package config

import (
	"log"
	"os"
	"strconv"
)

var (
	APP_ID        int
	API_HASH      string
	USER_SESSION  string
	BOT_TOKEN     string
	ChatList      = []int{
		-1001981472640,
		-1002400743054,
		-1002473285611,
		-1002340280965,
	}
)

func LoadConfig() {
	var err error
	
	appIDStr := os.Getenv("APP_ID")
	if appIDStr == "" {
		log.Fatal("APP_ID environment variable is required")
	}
	APP_ID, err = strconv.Atoi(appIDStr)
	if err != nil {
		log.Fatalf("Invalid APP_ID: %v", err)
	}

	API_HASH = os.Getenv("API_HASH")
	if API_HASH == "" {
		log.Fatal("API_HASH environment variable is required")
	}

	USER_SESSION = os.Getenv("USER_SESSION")
	if USER_SESSION == "" {
		log.Fatal("USER_SESSION environment variable is required")
	}

	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	if BOT_TOKEN == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}
}