package dao

import (
	"fmt"
)

const (
	keyCode = "code:%s"
)

// CodeSet sets code to cache
func (d *Dao) CodeSet(addr, code string) {
	d.Cache.Set(fmt.Sprintf(keyCode, addr), code, 0)
}

// CodeEqual code equals
func (d *Dao) CodeEqual(addr, code string) bool {
	if value := d.Cache.Get(fmt.Sprintf(keyCode, addr)); value != nil {
		return value == code
	}
	return false
}
