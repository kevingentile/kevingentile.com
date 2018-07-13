package main

import (
	"encoding/json"
	"net/http"
	"os"

	tracker "github.com/LaughingCabbage/fortnite-tracker/v1"
)

// Data holds response data to serve as json
type Data struct {
	Wins  int     `json:"wins"`
	KDR   float64 `json:"kdr"`
	Kills int     `json:"kills"`
}

// TODO rate limit this handler
func handleFortniteData(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("KEY")
	profile, err := tracker.GetProfile("pc", "laughingcabbage", key)
	if err != nil {
		handleError(err, w)
	}
	data := Data{}

	kills, err := tracker.GetKills(profile)
	handleError(err, w)
	data.Kills = kills

	wins, err := tracker.GetWins(profile)
	handleError(err, w)
	data.Wins = wins

	kdr, err := tracker.GetCurrentKDR(profile)
	handleError(err, w)
	data.KDR = kdr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		data := Data{
			Wins:  -1,
			KDR:   -10.0,
			Kills: 2,
		}
		json.NewEncoder(w).Encode(data)
	}
}
