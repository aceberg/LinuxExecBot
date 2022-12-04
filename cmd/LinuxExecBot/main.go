package main

import (
	"flag"
	"log"

	"github.com/aceberg/LinuxExecBot/internal/bot"
	"github.com/aceberg/LinuxExecBot/internal/yaml"
)

const confPath = "config.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	data := yaml.Read(*confPtr)

	if data.Conf.Token == "" || data.Conf.ID == 0 {
		log.Printf("ERROR: no Token or Chat ID")
		return
	}

	if data.Coms == nil {
		log.Printf("ERROR: no commands")
		return
	}

	bot.Start(data)
}
