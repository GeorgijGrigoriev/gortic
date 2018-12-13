package server

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"../config"
	"github.com/georgijgrigoriev/gortic/models"
	"github.com/gorilla/mux"

	//Go Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Run - main server instance
func Run(cfg *config.Config) {
	message := `Hello, this is Go Ticket's System Version 3
		Unstable version 
		Refer to github.com/georgijgrigoriev/gortic/ for any help
		Enjoy :)`
	log.Println(message)
	CreateFolderIfNotExist("logs")
	log.Printf("Starting server on %s\n", cfg.ListenSpec)

	router := mux.NewRouter()
	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	router.NotFoundHandler = NotFound404
	router.PathPrefix("/assets/").Handler(assets)
	router.HandleFunc("/tickets/", showTickets)
	router.HandleFunc("/archive/", showArchive)
	router.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.ListenSpec,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

//indexHandler - handler of index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//Logging remote address
	RenderTemplate(w, "index.html")
}

//RequestLogging - logs requests
func RequestLogging(r *http.Request) {
	err := os.Chdir("logs")
	Check(err)
	name := "req.log"
	wl, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0600)
	Check(err)
	_, err = wl.WriteString(time.Now().String() + " : " + r.RemoteAddr + " : " + r.RequestURI + "\n")
	err = wl.Sync()
	Check(err)
	defer wl.Close()
}

//NotFound404 - 404 error handler
var NotFound404 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "404.html")
})

var showTickets = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	data := models.GetTickets()
	fmt.Fprint(w, data.Sta)
})

var showArchive = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "archive.html")
})

//WaitForSignalTerm - counter + exit program
func WaitForSignalTerm() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	s := <-ch
	if s == os.Interrupt {
		for i := 2; i > 0; i-- {
			log.Printf("Terminating app in %v seconds", i)
			time.Sleep(time.Second)
		}
		os.Exit(3)
	} else {
		log.Println(s)
	}

}

//Health - Check server health
func Health() {
	log.Println("In dev")
}

//RenderTemplate - render template function
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateDir := filepath.Join("templates", tmpl)
	t := template.Must(template.ParseFiles(templateDir))
	t.ExecuteTemplate(w, tmpl, nil)
}

//Check - error handler
func Check(e error) {
	if e != nil {
		log.Println(e)
	}
}

//CreateFolderIfNotExist - creating folder if not exists
func CreateFolderIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0777)
		Check(err)
	}
}

//Open - sql open connection
func Open(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBType, cfg.DBConn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	return db, nil
}

func getTickets(db *Open) {

}
