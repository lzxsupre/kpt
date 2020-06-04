package model

// ArgScanRec arg
type ArgScanRec struct {
	UID string  `json:"uid" validate:"required"`
	CID string  `json:"cid" validate:"required"`
	Tpt float64 `json:"tpt" validate:"required"`
}

// ArgDateBetween arg
type ArgDateBetween struct {
	From string `form:"from" validate:"required"`
	To   string `form:"to" validate:"required"`
}

// ArgAuth arg
type ArgAuth struct {
	UID string `form:"uid"`
}
