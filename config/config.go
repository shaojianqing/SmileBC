package config

import (
	"encoding/json"
	"os"
)

func LoadConfigFromFile(filepath string) (*Config, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

type Config struct {
	DBConfig   DBConfig   `json:"dbConfig"`
	HttpConfig HttpConfig `json:"httpConfig"`
	SyncConfig SyncConfig `json:"syncConfig"`
}

type DBConfig struct {
	DBFilePath string `json:"dbFilePath"`
	CacheSize  int    `json:"cacheSize"`
	Handlers   int    `json:"handlers"`
}

type HttpConfig struct {
	ServerPort string `json:"serverPort"`
}

type SyncConfig struct {
}
