package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// AddUser inserts
func (d *Dao) AddUser(c context.Context, user *model.User) error {
	return d.DB.Create(user).Error
}

// UpdateUser updates
func (d *Dao) UpdateUser(c context.Context, user *model.User) error {
	return d.DB.Model(&model.User{}).Update(user).Where("uid=?", user.UID).Error
}

// QueryUsers selects
func (d *Dao) QueryUsers(c context.Context, user *model.User) (res *model.UserResponse, err error) {
	res = &model.UserResponse{}
	err = d.DB.Model(&model.User{}).Where(user).Count(&res.Total).Find(&res.Users).Error
	return
}

// DeleteUser deletes user
func (d *Dao) DeleteUser(c context.Context, uid string) error {
	return d.DB.Model(&model.User{}).Where("UID=?", uid).Update("status", 0).Error
}
