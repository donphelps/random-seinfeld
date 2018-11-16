package main

import (
	"net/http"
	"random-seinfeld/episodes"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func randomAPIHandler(w http.ResponseWriter, r *http.Request) {
	var e episodes.Episode

	vars := mux.Vars(r)
	s := vars["selector"]

	switch s {
	case "jerry":
		e = *jerryList.Random()
	case "george":
		e = *georgeList.Random()
	case "kramer":
		e = *kramerList.Random()
	case "elaine":
		e = *elaineList.Random()
	default:
		e = *fullList.Random()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(e.AsJSONbytes())
}
