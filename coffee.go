package main

import (
	"golib/logs"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"telegram_bot/register"
	"telegram_bot/events"
)

func main() {

	defer logs.Recover()

	bot, err := tgbotapi.NewBotAPI("413192403:AAEmBRb-nkSmZCkUUPt5bUD9PB0mVWJq22Q")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:

			if update.Message == nil {
				continue
			}

			var msgText string

			if register.IsRegistered(update.Message.Chat.ID) {

				switch update.Message.Text {
				case "баланс", "Баланс", "бал", "Бал", "БАЛАНС", "БАЛ", "Balance", "balance", "Bal", "bal":

					msgText = events.GetBalance(update.Message.Chat.ID)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)

					bot.Send(msg)

				case "Помощь", "помощь", "Help", "help", "support", "Support":

					msgText = events.GetHelpMg(update.Message.Chat.ID)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
					bot.Send(msg)

				case "photo":

					photo := events.GetPhoto(update.Message.Chat.ID)

					stickerconf := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photo)
					bot.Send(stickerconf)

				}
			} else if update.Message.Text == "/start" {
				register.SaveUser(update.Message.Chat.ID)
				msgText = events.GetWelcomeMsg()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				bot.Send(msg)

			} else {
				msgText = events.GetGreetingMsg()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				bot.Send(msg)
			}
		}
	}
}
