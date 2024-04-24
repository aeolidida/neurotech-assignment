package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	FileStorage `yaml:"file_storage"`
	HTTPServer  `yaml:"server"`
}

type FileStorage struct {
	Path        string `yaml:"path"`
	IsEncrypted bool   `yaml:"is_encrypted"`
	Secret      string `yaml:"secret" env:"STORAGE_SECRET_KEY"`
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func Init() (*Config, error) {
	// Попытка считать путь файла с конфигами
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("config.Init CONFIG_PATH is not set")
	}

	// Проверка существует ли файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config.Init config file does not exist: %s", configPath)
	}

	// Чтение конфига
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("config.Init cannot read config: %s", err)
	}

	return &cfg, nil
}
