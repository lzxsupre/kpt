package conf

import (
	"github.com/mivinci/abc/services/email"
	"github.com/mivinci/aladin"
	"github.com/mivinci/kpt/dao"
)

// Config is root config
type Config struct {
	DB    *dao.Config
	Key   *Key
	Bus   *Bus
	Email *email.Config
}

// New new
func New() *Config {
	c := &Config{}
	aladin.Init()
	aladin.Watch("cmd/db.yml", &c.DB)
	aladin.Watch("cmd/key.yml", &c.Key)
	aladin.Watch("cmd/email.yml", &c.Email)
	aladin.Watch("cmd/bus.yml", &c.Bus)
	return c
}
