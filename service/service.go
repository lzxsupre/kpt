package service

import (
	"html/template"

	"github.com/mivinci/abc/services/email"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/dao"
)

// Service service
type Service struct {
	mailer *email.Email
	tmpl   *template.Template
	dao    *dao.Dao
	c      *conf.Config
}

// New new
func New(c *conf.Config) *Service {
	return &Service{
		mailer: email.New(c.Email),
		tmpl:   newTmpl("cmd/warn.html"),
		dao:    dao.New(c.DB),
		c:      c,
	}
}

func newTmpl(path string) *template.Template {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}
	return tmpl
}

// Close closes
func (s *Service) Close() error {
	return s.dao.Close()
}
