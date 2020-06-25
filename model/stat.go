package model

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
