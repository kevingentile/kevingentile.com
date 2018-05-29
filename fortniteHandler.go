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
		panic(err)
	}
	data := Data{}

	kills, err := tracker.GetKills(profile)
	handleError(err)
	data.Kills = kills

	wins, err := tracker.GetWins(profile)
	handleError(err)
	data.Wins = wins

	kdr, err := tracker.GetKDR(profile)
	handleError(err)
	data.KDR = kdr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
