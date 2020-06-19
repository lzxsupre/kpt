package model

import "github.com/mivinci/abc/time"

// TempRec temperature record
type TempRec struct {
	ID    int       `json:"id"`
	UID   string    `json:"uid" validate:"required"`
	Temp  float32   `json:"temp" validate:"required"`
	Ctime time.Time `json:"ctime" gorm:"-"`
}

// TableName table name
func (TempRec) TableName() string {
	return "temp_record"
}
