package model

import "github.com/mivinci/abc/time"

// User user
type User struct {
	ID      int64     `json:"id" form:"id" gorm:"primary_key"`
	UID     string    `json:"uid" form:"uid"`
	Cid     string    `json:"cid" form:"cid"`
	ClassID string    `json:"class_id" form:"class_id"`
	Name    string    `json:"name" form:"name"`
	Email   string    `json:"email" form:"email"`
	Status  int8      `json:"status" form:"status"`
	Ctime   time.Time `json:"ctime" gorm:"-"`
	Mtime   time.Time `json:"mtime" gorm:"-"`
}

// TableName table name
func (User) TableName() string {
	return "user"
}

// UserResponse user response
type UserResponse struct {
	Total int     `json:"total"`
	Users []*User `json:"users"`
}
