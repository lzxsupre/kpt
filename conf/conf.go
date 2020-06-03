package conf

import (
	"github.com/mivinci/aladin"
	"github.com/mivinci/kpt/dao"
)

// Config is root config
type Config struct {
	DB *dao.Config
}

// Global Configs
var (
	DBConfig = &dao.Config{}
)

// New new
func New() *Config {
	aladin.Init()
	aladin.Watch("cmd/db.yml", DBConfig)

	return &Config{
		DB: DBConfig,
	}
}
