package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Addr        string `yaml:"addr"`
	DBPath      string `yaml:"db_path"`
	LoggerLevel string `yaml:"logger_level"`
	Format      string `yaml:"logger_format"`
}

func New(path string) *Config {
	if path == "" {
		path = "config.yml"
	}
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic("config err " + err.Error())
	}
	return cfg
}
