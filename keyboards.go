package main

import (
	"github.com/crimesbot/bot/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendKeyboard(u *session.Session, text string) {

	message := tgbotapi.NewMessage(u.ChatID, text)

	switch u.State[len(u.State)] {
	case 1:

	case 2:

		btnPrevious := tgbotapi.NewKeyboardButton("Previous")
		btnNewer := tgbotapi.NewKeyboardButton("Next")
		btnRow1 := tgbotapi.NewKeyboardButtonRow(btnPrevious, btnNewer)
		btnAdd := tgbotapi.NewKeyboardButton(string('\U0001F44D'))
		btnLove := tgbotapi.NewKeyboardButton(string('\U00002764'))
		btnSub := tgbotapi.NewKeyboardButton(string('\U0001F44E'))
		btnRow2 := tgbotapi.NewKeyboardButtonRow(btnAdd, btnLove, btnSub)
		keyboard := tgbotapi.NewReplyKeyboard(btnRow1, btnRow2)

		keyboard.OneTimeKeyboard = false

		message.ReplyMarkup = keyboard
		_, err := u.Send(message)
		if err != nil {
			u.Println("Error sending Config keyboard:" + err.Error())
		}
	}
}

func SendRateInline(u *session.Session) {

	u.Println("It's an opensource project, help us!")

	u.Println("Please access the link below and give us five Stars!")

	message := tgbotapi.NewMessage(u.ChatID, "https://telegram.me/storebot?start=textpicbot")
	message.DisableWebPagePreview = true

	_, err := u.Send(message)
	if err != nil {
		u.Println("Error sending Rate keyboard: " + err.Error())
	}

}
