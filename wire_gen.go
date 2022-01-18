// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/AidenHadisi/MyDailyBibleBot/api"
	"github.com/AidenHadisi/MyDailyBibleBot/bot"
	"github.com/AidenHadisi/MyDailyBibleBot/config"
)

// Injectors from wire.go:

func InitializeBot(cfg *config.Config) *bot.Bot {
	twitterApi := api.NewTwitterApi(cfg)
	client := _wireClientValue
	cache := _wireCacheValue
	bibleAPI := api.NewBibleAPI(client, cache)
	botBot := bot.NewBot(cfg, twitterApi, bibleAPI)
	return botBot
}

var (
	_wireClientValue = api.DefaultHttpClient
	_wireCacheValue  = api.DefaultCache
)
