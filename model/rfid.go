package model

import "github.com/mivinci/abc/time"

// RFIDRec rfid record
type RFIDRec struct {
	ID    int       `json:"id"`
	UID   string    `json:"uid" form:"uid" validate:"required"`
	Rfid  string    `json:"rfid" form:"rfid"`
	Type  int8      `json:"type" form:"type"`
	Ctime time.Time `json:"ctime" gorm:"-"`
}

// TableName table name
func (RFIDRec) TableName() string {
	return "rfid_record"
}
