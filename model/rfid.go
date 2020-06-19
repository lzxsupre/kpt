package model

import "github.com/mivinci/abc/time"

// RFIDRec rfid record
type RFIDRec struct {
	ID    int       `json:"id"`
	UID   string    `json:"uid"`
	Rfid  string    `json:"rfid"`
	Type  int8      `json:"type"`
	Ctime time.Time `json:"ctime" gorm:"-"`
}

// TableName table name
func (RFIDRec) TableName() string {
	return "rfid_record"
}
