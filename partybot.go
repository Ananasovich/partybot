package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGTOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	questions := readFile()
	text := "Подними руку и помаши или поищи глазами человека, который уже поднял руку. Обсудите вопрос. Начни сначала прислав мне любое сообщение. Удачи :) \n"

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text+sampleQuestion(questions))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func readFile() []string {
	data, _ := ioutil.ReadFile("Questions.txt")
	return strings.Split(string(data), "\n")
}

func sampleQuestion(questions []string) string {
	number := rand.Intn(len(questions) - 1)
	return questions[number]
}
