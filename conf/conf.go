package conf

import (
	"github.com/mivinci/aladin"
	"github.com/mivinci/kpt/dao"
)

// Key key
type Key struct {
	Secret string
}

// Config is root config
type Config struct {
	DB  *dao.Config
	Key *Key
}

// // Global Configs
// var (
// 	DBConfig  = &dao.Config{}
// 	KeyConfig = &Key{}
// )

// New new
func New() *Config {
	c := &Config{}
	aladin.Init()
	aladin.Watch("cmd/db.yml", &c.DB)
	aladin.Watch("cmd/key.yml", &c.Key)
	return c
}
