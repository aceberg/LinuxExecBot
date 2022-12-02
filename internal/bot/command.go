package bot

import (
	"log"
	"os/exec"

	"github.com/aceberg/LinuxExecBot/internal/models"
	"github.com/aceberg/LinuxExecBot/internal/check"
)

func execCommand(command string, allowedComm []models.Command) string {

	log.Println("INFO: received command:", command)

	out := "Unknown command"
	for _, oneCommand := range allowedComm {
		if oneCommand.Name == command {
			cmd, err := exec.Command(oneCommand.Exec).Output()
			check.IfError(err)

			out = string(cmd)
			break
		}
	}

	return out
}
