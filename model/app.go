package model

import "time"

//App app
type App struct {
	AppID  string    `json:"appid" form:"appkey"`
	AppKey string    `json:"appkey,omitempty" form:"appkey"`
	UID    string    `json:"uid" form:"uid"`
	Ctime  time.Time `json:"ctime" gorm:"-"`
}

//TableName table name
func (App) TableName() string {
	return "app"

}
