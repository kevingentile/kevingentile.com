package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	tracker "github.com/kevingentile/fortnite-tracker"
)

// Data holds response data to serve as json
type Data struct {
	Wins  int     `json:"wins"`
	KDR   float64 `json:"kdr"`
	Kills int     `json:"kills"`
}

func handleFortniteData(c *gin.Context) {
	platform := c.Param("platform")
	username := c.Param("username")
	key := viper.GetString("fortnite_tracker_token")
	profile, err := tracker.GetProfile(platform, username, key)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data := Data{}

	kills, err := profile.GetKills()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data.Kills = kills

	wins, err := profile.GetWins()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data.Wins = wins

	kdr, err := profile.GetCurrentKDR()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	data.KDR = kdr

	c.JSON(http.StatusOK, data)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println("Error proxying fornite stat request: ", err)
		data := Data{
			Wins:  -1,
			KDR:   -1,
			Kills: -1,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
