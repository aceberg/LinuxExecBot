package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/aceberg/LinuxExecBot/internal/models"
)

// menu - create bot menu
func menu(data models.Data) []tgbotapi.BotCommand {

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

	return allComm
}
