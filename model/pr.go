package model

import "github.com/mivinci/abc/time"

// PunchRec puch record
type PunchRec struct {
	ID                uint      `json:"id" gorm:"auto_increment,primary_key"`
	UID               string    `json:"uid" form:"uid" validate:"required"`
	Name              string    `json:"name" form:"name" validate:"required"`
	Phone             string    `json:"phone" form:"phone" validate:"required"`
	Location          string    `json:"location" form:"location" validate:"required"`
	IsTemperatureOK   bool      `json:"is_temperature_ok" form:"is_temperature_ok"`
	DidMeetHubei      bool      `json:"did_meet_hubei" form:"did_meet_hubei"`
	HasSymptom        bool      `json:"has_symptom" form:"has_symptom"`
	IsFamilyDiagnosed bool      `json:"is_family_diagnosed" form:"is_family_diagnosed"`
	DidMeetDiagnoses  bool      `json:"did_meet_diagnoses" form:"did_meet_diagnoses"`
	IsFamilySuspected bool      `json:"is_family_suspected" form:"is_family_suspected"`
	Ctime             time.Time `json:"ctime" gorm:"-"`
}

// TableName name
func (PunchRec) TableName() string {
	return "punch_record"
}
