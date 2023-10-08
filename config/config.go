package config

func LoadConfigFromFile(filepath string) *Config {
	return nil
}

type Config struct {
	DBConfig   DBConfig   `json:"dbConfig"`
	HttpConfig HttpConfig `json:"httpConfig"`
	SyncConfig SyncConfig `json:"syncConfig"`
}

type DBConfig struct {
	dbFilePath string `json:"dbFilePath"`
}

type HttpConfig struct {
	ServerPort string
}

type SyncConfig struct {
}
