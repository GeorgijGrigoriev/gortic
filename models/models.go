package models

import "github.com/georgijgrigoriev/gortic/db"

//GetTickets - get tickets from db
func GetTickets() {
	driver, err := db.Open()
	server.Check(err)
	data, err = driver.Prepare("Select * from tickets")
	defer data.Close()
	return data
}
