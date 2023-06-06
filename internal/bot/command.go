package bot

import (
	"log"
	"os/exec"

	"github.com/aceberg/LinuxExecBot/internal/check"
	"github.com/aceberg/LinuxExecBot/internal/models"
)

func execCommand(command, args string, data models.Data) string {

	log.Println("INFO: received command:", command)
	log.Println("INFO: received arguments:", args)

	out := "Unknown command"
	for _, oneCommand := range data.Coms {
		if oneCommand.Name == command && oneCommand.Exec != "" {

			var err error
			var cmd []byte
			if data.Conf.Args == "yes" {
				log.Println("INFO: executing", oneCommand.Exec, args)
				cmd, err = exec.Command("sh", "-c", oneCommand.Exec+" "+args).Output()
				check.IfError(err)
			} else {
				log.Println("INFO: executing", oneCommand.Exec)
				cmd, err = exec.Command("sh", "-c", oneCommand.Exec).Output()
				check.IfError(err)
			}

			out = string(cmd)
			break
		}
		if command == "help" {
			if out == "Unknown command" {
				out = "Available commands: \n"
			}

			out = out + "/" + oneCommand.Name + " - " + oneCommand.Desc + "\n"
		}
	}

	return out
}
