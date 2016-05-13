package aux

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

func GetUserID(update tgbotapi.Update) (int, error) {
	if update.Message != nil {
		return update.Message.From.ID, nil
	}
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID, nil
	}
	if update.InlineQuery != nil {
		return update.InlineQuery.From.ID, nil
	}
	if update.ChosenInlineResult != nil {
		return update.ChosenInlineResult.From.ID, nil
	}
	return 0, errors.New("UserID not found")
}

func GetChatID(update tgbotapi.Update) (int64, error) {

	if update.Message != nil {
		return update.Message.Chat.ID, nil
	}
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID, nil
	}
	if update.InlineQuery != nil {
		return 0, errors.New("ChatID not aplicable for InlineQuery")
	}
	if update.ChosenInlineResult != nil {
		return 0, errors.New("ChatID not aplicable for ChosenInlineResult")
	}
	return 0, errors.New("ChatID not found")
}
