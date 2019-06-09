# Fortnite Tracker API

[![CircleCI](https://circleci.com/gh/kevingentile/fortnite-tracker.svg?style=svg)](https://circleci.com/gh/kevingentile/fortnite-tracker)
[![GoDoc](https://godoc.org/github.com/kevingentile/fortnite-tracker?status.svg)](https://godoc.org/github.com/kevingentile/fortnite-tracker)
[![Release](https://img.shields.io/github/release/kevingentile/fortnite-tracker/all.svg)](https://github.com/kevingentile/fortnite-tracker/releases)

This project allows the user to make requests to the [FortniteTracker API](https://fortnitetracker.com/site-api) with a valid API token.

### Author: Kevin Gentile
### Contact: kevin@kevingentile.com

## Installation
```
go get github.com/kevingentile/fortnite-tracker
```

## Example Usage V2
```
	import "tracker github.com/kevingentile/fortnite-tracker"

	type Data struct {
		Wins  int     `json:"wins"`
		KDR   float64 `json:"kdr"`
		Kills int     `json:"kills"`
	}

	platform := "pc"
	username := "laughingcabbage"
	apiToken := os.Getenv("KEY")

	profile, err := tracker.GetProfile(platform, username, apiToken)
	if err != nil {
		handleError(err, w)
		return
	}

	kills, err := profile.GetKills()
	if err != nil {
		handleError(err, w)
		return
	}
	data.Kills = kills

	wins, err := profile.GetWins()
	if err != nil {
		handleError(err, w)
		return
	}
	data.Wins = wins

	kdr, err := profile.GetCurrentKDR()
	if err != nil {
		handleError(err, w)
		return
	}
	data.KDR = kdr

	fmt.Println("My Stats: ", data)
}
```


## Example Usage V1
```
	import "tracker github.com/kevingentile/fortnite-tracker"

	type Data struct {
		Wins  int     `json:"wins"`
		KDR   float64 `json:"kdr"`
		Kills int     `json:"kills"`
	}

	platform := "pc"
	username := "laughingcabbage"
	apiToken := os.Getenv("KEY")

	profile, err := tracker.GetProfile(platform, username, apiToken)
	if err != nil {
		handleError(err, w)
		return
	}

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

	fmt.Println("My Stats: ", data)
}
```