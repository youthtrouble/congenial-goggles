package telegram

import (
	"errors"
	"log"
	"os"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/youthtrouble/congenial-goggles/gpt"
)

func InitTelegramListening() {
	bot, err := telegrambot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Print(errors.New("Error initializing telegram bot: " + err.Error()))
	}

	bot.Debug = true
	u := telegrambot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		promptResponse, err := gpt.RetrieveOpenAIChatCompletions(update.Message.Text)
		if err != nil {
			log.Println(err)
			continue
		}

		msg := telegrambot.NewMessage(update.Message.Chat.ID, *promptResponse)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}