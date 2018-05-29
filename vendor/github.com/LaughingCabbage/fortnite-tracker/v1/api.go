package tracker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

// GetProfile is used to request a profile from fortnite tracker
func GetProfile(platform, name, APIToken string) (Profile, error) {
	profile := Profile{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", ProfileRoute+platform+"/"+name, nil)
	if err != nil {
		log.Fatal("failed to create profile request", err)
		return profile, err
	}
	req.Header.Add(headerKeyName, APIToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("failed to get profile from server", err)
		return profile, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &profile)
	if err != nil {
		log.Fatal("failed to unmarshal profile json")
		return profile, err
	}
	return profile, nil
}

// GetWins returns a profile lifetime wins
func GetWins(profile Profile) (int, error) {
	winsString, err := lookupLifetimeStat(profile, "Wins")
	if err != nil {
		return -1, err
	}
	wins, err := strconv.Atoi(winsString)
	if err != nil {
		return -1, err
	}
	return wins, nil
}

// GetTop3s returns a profile lifetime top 3 placements
func GetTop3s(profile Profile) (int, error) {
	top3String, err := lookupLifetimeStat(profile, "Top 3s")
	if err != nil {
		return -1, err
	}
	top3, err := strconv.Atoi(top3String)
	if err != nil {
		return -1, err
	}
	return top3, nil
}

// GetKills returns a profile lifetime kills
func GetKills(profile Profile) (int, error) {
	killsString, err := lookupLifetimeStat(profile, "Kills")
	if err != nil {
		return -1, err
	}

	kills, err := strconv.Atoi(killsString)
	if err != nil {
		return -1, err
	}

	return kills, nil
}

// GetKDR returns a profile lifetime kill death ratio
func GetKDR(profile Profile) (float64, error) {
	kdrString, err := lookupLifetimeStat(profile, "K/d")
	if err != nil {
		return -1, err
	}

	kdr, err := strconv.ParseFloat(kdrString, 64)
	if err != nil {
		return -1, err
	}
	return kdr, nil
}

func lookupLifetimeStat(profile Profile, key string) (string, error) {
	for _, item := range profile.LifeTimeStats {
		if item.Key == key {
			return item.Value, nil
		}
	}
	return "", ErrNotFound
}

// GetCurrentKDR returns the kdr for the current season
func GetCurrentKDR(profile Profile) (float64, error) {
	stats := profile.Stats
	kdrAverage := (stats.CurrP10.KD.ValueDec +
		stats.CurrP2.KD.ValueDec +
		stats.CurrP9.KD.ValueDec) / 3
	kdr := math.Round(kdrAverage*100) / 100
	return kdr, nil
}
