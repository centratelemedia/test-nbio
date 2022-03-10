package main

type Config struct {
	DBHost     string
	DbUser     string
	DbPassword string
	DBName     string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	//
}
