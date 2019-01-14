// Copyright (C) 2019 Ali Shanaakh, github@shanaakh.pro
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package bot

import (
	"github.com/shal/robot/pkg/util"
	"log"
)

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shal/robot/internal/app"
)

func Run() {
	jsonPath := util.GetEnv("BOT_DATA_PATH", "/tmp/bot.json")

	botApp := app.NewApp(jsonPath)

	// Bot updates configuration.
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := botApp.Bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Println(err.Error())
	}

	// Triggers, when new update pushed to channel.
	for update := range updates {
		botApp.ProcessUpdate(update)
	}
}
