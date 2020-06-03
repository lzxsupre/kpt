package model

// ScanRec scan
type ScanRec struct {
	UID   string   `json:"uid"`
	CID   string   `json:"cid"`
	Tpt   string   `json:"tpt"`
	Ctime datetime `json:"ctime"`
}

type datetime string
