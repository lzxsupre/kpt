package service

import (
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/dao"
)

// Service service
type Service struct {
	dao *dao.Dao
}

// New new
func New(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c.DB),
	}
}

// Close closes
func (s *Service) Close() error {
	return s.dao.Close()
}
