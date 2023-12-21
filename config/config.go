package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"

	"github.com/shaojianqing/smilebc/crypto"
)

const (
	MinimumSeedCount = 2
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

func ValidateConfiguration(config *Config) error {
	if config == nil {
		return fmt.Errorf("config does not exist")
	}

	if len(config.CommonConfig.PrivateKeyValue) == 0 {
		return fmt.Errorf("private key is not set yet")
	}

	privateKey, err := crypto.HexToECDSA(config.CommonConfig.PrivateKeyValue)
	if err != nil {
		return fmt.Errorf("private key is not correct, please check the value")
	}
	config.CommonConfig.PrivateKey = privateKey

	if len(config.DBConfig.DBFilePath) == 0 {
		return fmt.Errorf("database path is not set yet")
	}

	if len(config.NetworkConfig.SeedNodes) < MinimumSeedCount {
		return fmt.Errorf("seed nodes is not set correctly")
	}
	return nil
}

type Config struct {
	CommonConfig  CommonConfig  `json:"commonConfig"`
	DBConfig      DBConfig      `json:"dbConfig"`
	HttpConfig    HttpConfig    `json:"httpConfig"`
	SyncConfig    SyncConfig    `json:"syncConfig"`
	NetworkConfig NetworkConfig `json:"networkConfig"`
}

type CommonConfig struct {
	PrivateKeyValue string `json:"privateKeyValue"`
	PrivateKey      *ecdsa.PrivateKey
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

type NetworkConfig struct {
	ListenAddress string   `mapstructure:"listenAddress"`
	TCPPort       uint16   `mapstructure:"tcpPort"`
	UDPPort       uint16   `mapstructure:"udpPort"`
	SeedNodes     []string `mapstructure:"seedNodes"`
}
