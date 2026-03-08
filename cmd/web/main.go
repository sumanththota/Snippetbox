package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	 "html/template"
	"os"
	"github.com/sumanththota/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"web:pass@/snippetbox?parseTime=true",
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

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// initlizing a new instance of applicaiton containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetModel{DB: db},
		templateCache: templateCache,

	}

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
