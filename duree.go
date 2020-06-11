package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

// Duree hold the complete app data
type Duree struct {
	listenAddr       string
	bookmarkFilepath string
	bookmarks        Bookmarks
}

// Run starts the duree app
func (d *Duree) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", d.indexHandler()).Methods("GET")
	r.HandleFunc("/save", d.saveHandler()).Methods("POST")

	staticBox := rice.MustFindBox("static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(staticBox.HTTPBox())))

	log.Print("Listening on ", d.listenAddr)
	log.Fatal(http.ListenAndServe(d.listenAddr, r))
}

func (d *Duree) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// reading the bookmarks
		bookmarks, err := d.read(d.bookmarkFilepath)
		if err != nil {
			log.Println(err.Error(), d.bookmarkFilepath)
			http.Error(w, err.Error(), 500)
			return
		}

		// find a rice.Box
		templateBox, err := rice.FindBox("templates")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Cannot read templates", 500)
			return
		}
		// get file contents as string
		templateString, err := templateBox.String("index.html")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Cannot read index.html", 500)
			return
		}

		// parse and execute the template
		tmplIndex, err := template.New("index").Parse(templateString)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), 500)
			return
		}
		tmplIndex.Execute(w, struct {
			Bookmarks []Bookmarks
		}{bookmarks})

	}

}

func (d *Duree) saveHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// decode json request
		decoder := json.NewDecoder(r.Body)
		var bookmarks []Bookmarks
		err := decoder.Decode(&bookmarks)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write bookmarks to file
		err = d.write(d.bookmarkFilepath, bookmarks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"success\": true, \"msg\": \"data_saved\"}"))
	}

}
