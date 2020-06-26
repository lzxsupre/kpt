package model

import "github.com/mivinci/abc/time"

//UIDTemp uid temp
type UIDTemp struct {
	UID  string  `json:"uid"`
	Temp float32 `json:"temp"`
}

//UIDTemps uid temp
type UIDTemps struct {
	Total int        `json:"total"`
	Temps []*UIDTemp `json:"temps"`
}

//StatTempRec stat
type StatTempRec struct {
	Total    int       `json:"total"`
	Normal   *UIDTemps `json:"normal"`
	Abnormal *UIDTemps `json:"abnormal"`
}

// StatUserNoCheckIn stat
type StatUserNoCheckIn struct {
	User         *User     `json:"user"`
	LastCheckout time.Time `json:"last_checkout"`
}

// StatLastRFIDRec stat
type StatLastRFIDRec struct {
	Total int                 `json:"total"`
	Out   map[string]*RFIDRec `json:"out"`
	In    map[string]*RFIDRec `json:"in"`
}
