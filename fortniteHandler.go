package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	tracker "github.com/kevingentile/fortnite-tracker/v1"
)

// Data holds response data to serve as json
type Data struct {
	Wins  int     `json:"wins"`
	KDR   float64 `json:"kdr"`
	Kills int     `json:"kills"`
}

func handleFortniteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := os.Getenv("KEY")
	profile, err := tracker.GetProfile(vars["platform"], vars["username"], key)
	if err != nil {
		handleError(err, w)
		return
	}
	data := Data{}

	kills, err := tracker.GetKills(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.Kills = kills

	wins, err := tracker.GetWins(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.Wins = wins

	kdr, err := tracker.GetCurrentKDR(profile)
	if err != nil {
		handleError(err, w)
		return
	}
	data.KDR = kdr

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		handleError(err, w)
		return
	}
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		data := Data{
			Wins:  -1,
			KDR:   -1,
			Kills: -1,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
