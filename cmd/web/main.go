package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}

func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"web:pass@tcp(localhost:3306)/snippetbox?parseTime=true",
		"MySQL data source name",
	)
	flag.Parse()	

	// level logging

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()


	// initlizing a new instance of applicaiton containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	//establish a router

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// to server the files into the ui

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//creating a new server struct for configuring to our server
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	


	// start the webserver
	infoLog.Printf("Starting the server on %s", *addr)
	//instead of using http.ListenAndServe() we are using the server struct to listen and serve the requests
	err = srv.ListenAndServe()
	// helps the application to terminate immedialty
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
	return nil, err
	}
	if err = db.Ping(); err != nil {
	return nil, err
	}
	return db, nil
	}
