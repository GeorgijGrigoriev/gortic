package db

import (
	"database/sql"
	"log"

	"../config"
	//Go Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Open - creating a database object
func Open(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBType, cfg.DBConn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	return db, nil
}
