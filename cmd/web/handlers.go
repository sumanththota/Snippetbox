package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//create a handler named home
// w provides methods to establish http response and sends it to the user
// r is a struct which holds the information about the current request

func(app *application) home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}


	files := []string{
		"ui/html/base.layout.html",
		"ui/html/home.page.html",
		"ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w,err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}

func(app *application) showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

func(app *application) createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a snippet"))
}
