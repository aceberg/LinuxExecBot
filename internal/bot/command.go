package bot

import (
	"log"
	"os/exec"

	"github.com/aceberg/LinuxExecBot/internal/check"
	"github.com/aceberg/LinuxExecBot/internal/models"
)

func execCommand(command string, allowedComm []models.Command) string {

	log.Println("INFO: received command:", command)

	out := "Unknown command"
	for _, oneCommand := range allowedComm {
		if oneCommand.Name == command {
			log.Println("INFO: executing", oneCommand.Exec)

			cmd, err := exec.Command("sh", "-c", oneCommand.Exec).Output()
			check.IfError(err)

			out = string(cmd)
			break
		}
	}

	return out
}
