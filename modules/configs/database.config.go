package configs

type DBConfig struct {
	configs string
}

func NewDBConfig(configs string) *DBConfig {
	return &DBConfig{configs: configs}
}
