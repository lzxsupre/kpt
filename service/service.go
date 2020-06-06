package service

import (
	"github.com/mivinci/abc/services/email"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/dao"
)

// Service service
type Service struct {
	mailer *email.Email
	dao    *dao.Dao
	c      *conf.Config
}

// New new
func New(c *conf.Config) *Service {
	return &Service{
		mailer: email.New(c.Email),
		dao:    dao.New(c.DB),
		c:      c,
	}
}

// Close closes
func (s *Service) Close() error {
	return s.dao.Close()
}
