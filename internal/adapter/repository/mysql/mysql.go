package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	User            string
	Passwd          string `env:"PASSWORD,unset"`
	Host            string `envDefault:"127.0.0.1"`
	Port            int    `envDefault:"3306"`
	DBName          string `env:"DB"`
	MaxIdleConns    int    `envDefault:"2"`
	MaxOpenConns    int    `envDefault:"50"`
	ConnMaxLifeTime string `envDefault:"1m"`
}

func NewDB(c *Config) (db *sql.DB, err error) {
	time.Local, _ = time.LoadLocation("Asia/Seoul")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.User, c.Passwd, c.Host, c.Port, c.DBName)
	db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	lifetime, err := time.ParseDuration(c.ConnMaxLifeTime)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(lifetime)
	return
}
