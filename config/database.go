package config

import (
	"database/sql"
	"fmt"
)

func NewDBConnection() (*sql.DB, error) {
	user := GetEnvWithDefault("DB_USER", "root")
	password := GetEnvWithDefault("DB_PASSWORD", "")
	host := GetEnvWithDefault("DB_HOST", "localhost")
	port := GetEnvWithDefault("DB_PORT", "3306")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ddd_sample?parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

