package model

// ArgScanRec arg
type ArgScanRec struct {
	UID string  `json:"uid" validate:"required"`
	Cid string  `json:"cid" validate:"required"`
	Tpt float64 `json:"tpt" validate:"required"`
}

// ArgDateBetween arg
type ArgDateBetween struct {
	From string `form:"from" validate:"required"`
	To   string `form:"to" validate:"required"`
}

// ArgAuth arg
type ArgAuth struct {
	Code string `form:"code"`
	UID  string `form:"uid"`
}

// ArgRFIDRec arg
type ArgRFIDRec struct {
	UID  string `json:"uid" form:"uid"`
	Rfid string `json:"rfid" form:"rfid"`
	Type int8   `json:"type" form:"type"`
}

// TableName table name
func (ArgRFIDRec) TableName() string {
	return "rfid_record"
}

// ArgTempRec arg
type ArgTempRec struct {
	UID string `form:"uid"`
}

// TableName table name
func (ArgTempRec) TableName() string {
	return "temp_record"
}
