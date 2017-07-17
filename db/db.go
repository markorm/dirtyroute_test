package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Config struct {
	Username	string
	Password	string
	Database	string
	Port 		string
}

// Open the database connection
var DB *sql.DB
func Open(c *Config) error {
	var err error
	DB, err = sql.Open("mysql", c.Username + ":" + c.Password + "@/" + c.Database + "?parseTime=true")
	if err != nil {
		fmt.Println("Failed to open database connection")
		return err
	}
	return err
}

func GetConnection() *sql.DB {
	return DB
}
