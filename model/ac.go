package model

import "github.com/mivinci/abc/time"

// ScanRec scan
type ScanRec struct {
	UID   string    `json:"uid"`
	CID   string    `json:"cid"`
	Tpt   string    `json:"tpt"`
	Ctime time.Time `json:"ctime"`
}
