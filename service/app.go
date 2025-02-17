package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/lithammer/shortuuid"
	"github.com/mivinci/kpt/model"
)

// App 获取应用
func (s *Service) App(c context.Context, id string) (*model.App, error) {
	return s.dao.QueryApp(c, id)
}

// Apps 获取应用
func (s *Service) Apps(c context.Context, app *model.App) ([]*model.App, error) {
	return s.dao.QueryAppsSecure(c, app)
}

// AddApp 添加应用
func (s *Service) AddApp(c context.Context, app *model.App) (*model.App, error) {
	origin := &model.App{
		AppID:  shortuuid.New(),
		AppKey: shortuuid.New(),
		UID:    app.UID,
	}
	m := sha256.New()
	m.Write([]byte(origin.AppKey))
	app.AppKey = hex.EncodeToString(m.Sum(nil))
	app.AppID = origin.AppID
	app.UID = origin.UID
	return origin, s.dao.AddApp(c, app)
}

// UpdateApp 更新应用
func (s *Service) UpdateApp(c context.Context, app *model.App) error {
	return s.dao.UpdateApp(c, app)
}
