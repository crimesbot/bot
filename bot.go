package main

import (
	"log"

	"github.com/crimesbot/bot/errs"

	"github.com/crimesbot/bot/aux"

	"github.com/crimesbot/bot/session"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(token string) *Bot {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicln(err)
	}

	bot.Debug = true

	log.Printf("Bot connected on account %s\n", bot.Self.UserName)

	return &Bot{bot: bot}
}

func (b *Bot) Run() {

	news := tgbotapi.NewUpdate(0)
	news.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(news)

	for update := range updates {
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		chatID, err := aux.GetChatID(update)
		errs.Catch(err, nil)

		uid, err := aux.GetUserID(update)
		errs.Catch(err, nil)

		s := session.GetSession(uid)
		if s == nil {
			s = session.NewSession(uid)
			s.State = append(s.State, 1) // SateInit
		}

		// Handle the actual command
		go b.Handle(s, update)
	}
}

func (b *Bot) Handle(s *session.Session, u tgbotapi.Update) {

}

func (b *Bot) Println(s *session.Session, line string) (tgbotapi.Message, error) {
	m := tgbotapi.NewMessage(s.ChatID, line)
	m.DisableWebPagePreview = true
	return b.Send(m)
}

func (b *Bot) HideKeyboard(s *session.Session, text string) (tgbotapi.Message, error) {
	m := tgbotapi.NewMessage(s.ChatID, text)
	m.ReplyMarkup = tgbotapi.NewHideKeyboard(true)
	return b.Send(m)
}

func (b *Bot) Send(msg tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	return b.bot.Send(msg)
}

func (b *Bot) EditMessage(s *session.Session, msg tgbotapi.Message, text string) (tgbotapi.Message, error) {
	edit := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:    s.ChatID,
			MessageID: msg.MessageID,
		},
		Text: text,
	}
	return b.bot.Send(edit)
}
