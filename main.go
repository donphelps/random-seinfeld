// Package Random-Seinfeld is a small web app for randomly selecting Seinfeld episodes to watch. Currently only configured to show episodes from Hulu.com. A subscription may be required.
package main

import (
	"html/template"
	"log"
	"net/http"
	"random-seinfeld/episodes"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const address string = "127.0.0.1:8000"

var (
	templates *template.Template

	fullList   episodes.EpisodeList
	jerryList  episodes.EpisodeList
	georgeList episodes.EpisodeList
	kramerList episodes.EpisodeList
	elaineList episodes.EpisodeList
)

func main() {
	loadEpisodeLists()
	loadTemplates()

	r := configureRoutes()
	s := configureServer(r)
	log.Fatal(startServer(s))
}

func loadEpisodeLists() {
	fullList.LoadFromJSON("assets/episodes.json")

	for i := range fullList {
		if strings.Contains(fullList[i].Description, "Jerry") {
			jerryList.Add(*fullList[i])
		}
		if strings.Contains(fullList[i].Description, "George") {
			georgeList.Add(*fullList[i])
		}
		if strings.Contains(fullList[i].Description, "Kramer") {
			kramerList.Add(*fullList[i])
		}
		if strings.Contains(fullList[i].Description, "Elaine") {
			elaineList.Add(*fullList[i])
		}
	}

	log.Println("Total episode count:", len(fullList))
	log.Println("Jerry episode count:", len(jerryList))
	log.Println("George episode count:", len(georgeList))
	log.Println("Kramer episode count:", len(kramerList))
	log.Println("Elaine episode count:", len(elaineList))
}

func loadTemplates() {
	templates = template.Must(template.ParseGlob("assets/html/*"))
}

func configureRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/rnd", randomAPIHandler)
	r.HandleFunc("/rnd/{selector}", randomAPIHandler)

	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	return r
}

func configureServer(r *mux.Router) *http.Server {
	return &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func startServer(s *http.Server) error {
	log.Printf("Starting listener at %v...\n", address)
	return s.ListenAndServe()
}
