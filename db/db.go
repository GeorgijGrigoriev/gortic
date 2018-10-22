package db

import (
	"database/sql"
	"log"

	"github.com/georgijgrigoriev/gortic/server"
	//Go Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Open - creating a database object
func Open(cfg *server.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBType, cfg.DBConn)
	if err != nil {
		log.Println(err)
	}

	return db, nil
}
