package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Config struct {
	DBUrl           string `env:"DB_URL"`
	MaxIdleConns    int    `envDefault:"2"`
	MaxOpenConns    int    `envDefault:"50"`
	ConnMaxLifeTime string `envDefault:"1m"`
}

func NewSqliteDB(c *Config) (db *sql.DB) {
	time.Local, _ = time.LoadLocation("Asia/Seoul")
	db, err := sql.Open("sqlite3", c.DBUrl)
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
