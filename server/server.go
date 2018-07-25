package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

//Config - main config struct
type Config struct {
	ListenSpec string
	DBConn     string
	DBType     string
	Assets     string
}

//Run - main server instance
func Run(cfg *Config) {
	log.Printf("Starting server on %s\n", cfg.ListenSpec)

	router := mux.NewRouter()
	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	router.NotFoundHandler = NotFound404
	router.PathPrefix("/assets/").Handler(assets)
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
	log.Println(r.RemoteAddr)
	RenderTemplate(w, "index.html")
}

//NotFound404 - 404 error handler
var NotFound404 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "404.html")
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

//Health - check server health
func Health() {
}

//RenderTemplate - render template function
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateDir := filepath.Join("templates", tmpl)
	t := template.Must(template.ParseFiles(templateDir))
	t.ExecuteTemplate(w, tmpl, nil)
}
