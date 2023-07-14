package models

// Config - telegram chat token and ID
type Config struct {
	Token string  `yaml:"token"`
	ID    int64   `yaml:"id"`
	IDs   []int64 `yaml:"ids"`
	Args  string  `yaml:"args"`
}

// Command - one command
type Command struct {
	Name string `yaml:"name"`
	Exec string `yaml:"exec"`
	Desc string `yaml:"desc"`
}

// Data - all from config file
type Data struct {
	Conf Config    `yaml:"config"`
	Coms []Command `yaml:"commands"`
}
