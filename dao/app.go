package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

//AddApp inserts
func (d *Dao) AddApp(c context.Context, app *model.App) error {
	return d.DB.Create(app).Error
}

//UpdateApp updates
func (d *Dao) UpdateApp(c context.Context, app *model.App) error {
	return d.DB.Model(&model.App{}).Update(app).Where("appid=?", app.AppID).Error
}

//QueryAppsSecure selects multiple
func (d *Dao) QueryAppsSecure(c context.Context, app *model.App) (apps []*model.App, err error) {
	return apps, d.DB.Model(&model.App{}).Select("app_id,uid,ctime").Where(app).Find(&apps).Error
}

//QueryApp selects
func (d *Dao) QueryApp(c context.Context, id string) (app *model.App, err error) {
	app = &model.App{}
	err = d.DB.Model(&model.App{}).Where("app_id=?",id).First(app).Error
	return
} 

