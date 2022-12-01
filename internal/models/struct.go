package models

// Config - telegram chat token and ID
type Config struct {
	Token string `yaml:"token"`
	ID    int64  `yaml:"id"`
}

// Command - one command
type Command struct {
	Name string `yaml:"name"`
	Exec string `yaml:"exec"`
}

// Data - all from config file
type Data struct {
	Conf Config    `yaml:"config"`
	Coms []Command `yaml:"commands"`
}
