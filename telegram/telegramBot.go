package telegram

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/youthtrouble/congenial-goggles/gpt"
	oandastuff "github.com/youthtrouble/congenial-goggles/oanda-stuff"
)

func InitAlfredTelegramListening() {
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

func InitOandaTelegramListening() {
	bot, err := telegrambot.NewBotAPI(os.Getenv("OANDA_BOT_TOKEN"))
	if err != nil {
		log.Print(errors.New("Error initializing telegram bot: " + err.Error()))
	}

	bot.Debug = true
	u := telegrambot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil || !strings.Contains(update.Message.Text, "@oandaabokibot") {
			continue
		}

		respoonseMessage := "Error getting updates, please try again in a bit"
		oandaRates, time, err := oandastuff.FetchCurrentOandaRates()
		if err == nil {
			respoonseMessage = fmt.Sprintf(" Current GBP/NGN rates: ₦%s\nTime: %s\n", oandaRates.Response[0].AverageBid, *time)
		}

		msg := telegrambot.NewMessage(getChatID(), respoonseMessage)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func getChatID() int64 {
	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		return 0
	}

	return int64(chatID)
}
