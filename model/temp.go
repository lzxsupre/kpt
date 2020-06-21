package model

import "github.com/mivinci/abc/time"

// TempRec temperature record
type TempRec struct {
	ID    int       `json:"id"`
	UID   string    `json:"uid" form:"uid" validate:"required"`
	Temp  float32   `json:"temp"`
	Ctime time.Time `json:"ctime" gorm:"-"`
}

// TableName table name
func (TempRec) TableName() string {
	return "temp_record"
}
