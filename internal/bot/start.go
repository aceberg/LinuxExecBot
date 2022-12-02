package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/aceberg/LinuxExecBot/internal/models"
	"github.com/aceberg/LinuxExecBot/internal/check"
)

// Start telegram bot
func Start(data models.Data) {

	bot, err := tgbotapi.NewBotAPI(data.Conf.Token)
	check.IfError(err)

	log.Printf("INFO: Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Chat.ID == data.Conf.ID { // Only from this Chat ID
			if update.Message == nil { // ignore any non-Message updates
				continue
			}
			if !update.Message.IsCommand() { // ignore any non-command Messages
				continue
			}

			msg := tgbotapi.NewMessage(data.Conf.ID, "")

			msg.Text = execCommand(update.Message.Command(), data.Coms)

			_, err := bot.Send(msg);
			check.IfError(err)
		}
	}
}
