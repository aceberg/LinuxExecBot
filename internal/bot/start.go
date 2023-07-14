package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/aceberg/LinuxExecBot/internal/check"
	"github.com/aceberg/LinuxExecBot/internal/models"
)

// Start telegram bot
func Start(data models.Data) {

	bot, err := tgbotapi.NewBotAPI(data.Conf.Token)
	check.IfError(err)

	log.Printf("INFO: Authorized on account %s", bot.Self.UserName)

	// Create commands Menu
	allComm := menu(data)

	// Delete old Menu
	cfgDel := tgbotapi.NewDeleteMyCommands()
	_, err = bot.Request(cfgDel)
	check.IfError(err)
	// Set Menu for each ID
	for _, id := range data.Conf.IDs {
		scope := tgbotapi.NewBotCommandScopeChat(id)
		cfgSet := tgbotapi.NewSetMyCommandsWithScope(scope, allComm...)
		_, err = bot.Request(cfgSet)
		check.IfError(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if check.InSlice(update.Message.Chat.ID, data.Conf.IDs) { // Only from this Chat ID

			if update.Message == nil { // ignore any non-Message updates
				continue
			}
			if !update.Message.IsCommand() { // ignore any non-command Messages
				continue
			}

			returnText := execCommand(update.Message.Command(), update.Message.CommandArguments(), data)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = returnText

			_, err = bot.Send(msg)
			check.IfError(err)
		}
	}
}
