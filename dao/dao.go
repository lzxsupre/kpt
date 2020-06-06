package dao

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mivinci/abc/cache"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Config config
type Config struct {
	DSN     string        `yml:"dsn"`
	Timeout time.Duration `yml:"timeout"`
}

// Dao dao
type Dao struct {
	Cache *cache.Cache
	DB    *gorm.DB
	db    *sql.DB
	c     *Config
}

// New new
func New(c *Config) *Dao {
	return &Dao{
		Cache: openCache(),
		DB:    openORM(c.DSN),
		db:    openDB(c.DSN),
		c:     c,
	}
}

func openDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func openORM(dsn string) *gorm.DB {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func openCache() *cache.Cache {
	return cache.New(time.Minute * 10)
}

// Close closes
func (d *Dao) Close() error {
	err1 := d.DB.Close()
	err2 := d.db.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}
