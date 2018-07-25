package main

import (
	"flag"

	"github.com/georgijgrigoriev/gortic/server"
)

func processFlags() *server.Config {
	cfg := &server.Config{}
	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:8080", "Setting address string and port of server")
	flag.StringVar(&cfg.DBConn, "db", "dbuser:dbpassword@tcp(192.168.9.22)/tickets", "Setting address,username,password of database (*username:*password@tcp(ip sql server)/database")
	flag.StringVar(&cfg.DBType, "dbtype", "mysql", "Setting what database is used (mysql default)")
	flag.StringVar(&cfg.Assets, "assets", "assets", "Setting where assets placed")
	flag.Parse()
	return cfg
}

func main() {
	cfg := processFlags()
	go server.WaitForSignalTerm()
	server.Run(cfg)
}
