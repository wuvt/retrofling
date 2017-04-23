package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) < 1 {
		http.Error(w, "No date specified", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("WUVTFM_20060102_1500Z", r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utcStartDate := time.Date(2017, 4, 23, 8, 0, 0, 0, time.UTC)
	if parsedDate.After(utcStartDate) {
		http.Redirect(w, r, fmt.Sprintf("https://archive.org/details/WUVTFM_%s00Z", parsedDate.Format("20060102_15")), http.StatusSeeOther)
	} else {
		easternTime, err := time.LoadLocation("America/New_York")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("https://archive.org/details/WUVTFM_%s", parsedDate.In(easternTime).Format("2006010215")), http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
