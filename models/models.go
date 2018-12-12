package models

import (
	"log"

	"../db"
	"../server"
)

//GetTickets - get tickets from db
func GetTickets(cfg *server.Config) {
	driver, err := db.Open(cfg.DBConn)
	server.Check(err)
	data, err := driver.Prepare("Select * from tickets")
	if err != nil {
		log.Println(err)
	}
	defer data.Close()
}
