package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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

	if len(config.CommonConfig.PrivateKey) == 0 {
		return fmt.Errorf("private key is not set yet")
	}

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
	P2PConfig     P2PConfig     `json:"p2pConfig"`
	SyncConfig    SyncConfig    `json:"syncConfig"`
	NetworkConfig NetworkConfig `json:"networkConfig"`
}

type CommonConfig struct {
	PrivateKey string `json:"privateKey"`
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

type P2PConfig struct {
	ListenAddress    string        `json:"listenAddress"`
	DialTimeout      time.Duration `json:"dialTimeout"`
	HandshakeTimeout string        `json:"handshakeTimeout"`
	MaxPeerCount     uint32        `json:"maxPeerCount"`
}

type NetworkConfig struct {
	ListenAddress string   `json:"listenAddress"`
	TCPPort       uint16   `json:"tcpPort"`
	UDPPort       uint16   `json:"udpPort"`
	SeedNodes     []string `json:"seedNodes"`
}
