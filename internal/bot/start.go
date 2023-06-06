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
	var oneComm tgbotapi.BotCommand
	var allComm []tgbotapi.BotCommand
	help := true
	for _, comm := range data.Coms {
		if comm.Name != "" && comm.Desc != "" {
			oneComm.Command = comm.Name
			oneComm.Description = comm.Desc

			allComm = append(allComm, oneComm)
		}
		if comm.Name == "help" {
			help = false
		}
	}
	if help {
		oneComm.Command = "help"
		oneComm.Description = "List all commands"
		allComm = append(allComm, oneComm)
	}

	cfg := tgbotapi.NewSetMyCommands(allComm...)
	_, err = bot.Request(cfg)
	check.IfError(err)

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

			msg.Text = execCommand(update.Message.Command(), update.Message.CommandArguments(), data)

			_, err = bot.Send(msg)
			check.IfError(err)
		}
	}
}
