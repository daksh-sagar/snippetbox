package main
import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr: *addr,
		ErrorLog: errLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	err := server.ListenAndServe()
	errLog.Fatal(err)
}