package repo

import (
	"database/sql"
	"fmt"
	"log"
)

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func InitPostgresConn() (*sql.DB, error) {

	connectionInfo := DBConfig{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "password",
		DB:       "ddgf_db_dev",
	}

	conn, err := sql.Open(connectionInfo.Driver, dbConfigToString(connectionInfo))
	if err != nil {
		log.Fatalf("Failed to open connection to database: %s\n", err.Error())
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s\n", err.Error())
		return nil, err
	}

	return conn, nil
}

func dbConfigToString(d DBConfig) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		d.Driver,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DB)
}
