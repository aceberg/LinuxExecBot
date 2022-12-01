package main

import (
	"flag"
	// "log"

	"github.com/aceberg/LinuxExecBot/internal/bot"
	"github.com/aceberg/LinuxExecBot/internal/yaml"
)

const confPath = "/data/LinuxExecBot/config.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	data := yaml.Read(*confPtr)

	// if data == models.Data{} {
	// 	log.Printf("ERROR: no data")
	// 	return
	// }

	bot.Start(data)
}
