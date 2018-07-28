// Copyright 2018 Kevin Gentile.
// Licensed under GNU General Public License v3.0
package tracker

// StatField is an abstract JSON object for a stats field
type StatField struct {
	Label        string  `json:"label"`
	Field        string  `json:"field"`
	Category     string  `json:"category"`
	ValueInt     int     `json:"valueInt"`
	Value        string  `json:"value"`
	Percentile   float64 `json:"percentile"`
	DisplayValue string  `json:"displayValue"`
}

// FloatStatField is an abstract JSON object for a decimal stat field
type FloatStatField struct {
	Label        string  `json:"label"`
	Field        string  `json:"field"`
	Category     string  `json:"category"`
	ValueDec     float64 `json:"valueDec"`
	Value        string  `json:"value"`
	Rank         int     `json:"rank"`
	Percentile   float64 `json:"percentile"`
	DisplayValue string  `json:"displayValue"`
}

// Profile json
type Profile struct {
	AccountID        string        `json:"accountId"`
	PlatformID       int           `json:"platformId"`
	PlatformName     string        `json:"platformName"`
	PlatformNameLong string        `json:"platformNameLong"`
	EpicUserHandle   string        `json:"epicUserHandle"`
	Stats            PersonalStats `json:"stats"`
	LifeTimeStats    LifeTimeStats `json:"lifeTimeStats"`
	RecentMatches    RecentMatches `json:"recentMatches"`
}

// Stat json
type Stat struct {
	TRNRating     StatField      `json:"trnRating"`
	Score         StatField      `json:"score"`
	Top1          StatField      `json:"top1"`
	Top3          StatField      `json:"top3"`
	Top5          StatField      `json:"top5"`
	Top6          StatField      `json:"top6"`
	Top10         StatField      `json:"top10"`
	Top12         StatField      `json:"top12"`
	Top25         StatField      `json:"top25"`
	KD            FloatStatField `json:"kd"`
	WinRatio      FloatStatField `json:"winRatio"`
	Matches       StatField      `json:"matches"`
	Kills         StatField      `json:"kills"`
	KPG           FloatStatField `json:"kpg"`
	ScorePerMatch FloatStatField `json:"scorePerMatch"`
}

// Personal
type PersonalStats struct {
	P2       Stat `json:"p2"`
	P10      Stat `json:"p10"`
	P9       Stat `json:"p9"`
	CurrP2   Stat `json:"curr_p2"`
	CurrP10  Stat `json:"curr_p10"`
	CurrP9   Stat `json:"curr_p9"`
	PriorP2  Stat `json:"prior_p2"`
	PriorP10 Stat `json:"prior_p10"`
	PriorP9  Stat `json:"prior_p9"`
}

// LifeTimeStats holds an array of lifetime stats
type LifeTimeStats []LifeTimeKey

// LifeTimeKey holds the key and value for lifetime stats map
type LifeTimeKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RecentMatches
type RecentMatches []Match

// Match
type Match struct {
	ID              int     `json:"id"`
	AcccountID      string  `json:"accountId"`
	Playlist        string  `json:"playlist"`
	Kills           int     `json:"kills"`
	MinutesPlayed   int     `json:"minutesPlayed"`
	Top1            int     `json:"top1"`
	Top5            int     `json:"top5"`
	Top6            int     `json:"top6"`
	Top10           int     `json:"top10"`
	Top12           int     `json:"top12"`
	Top25           int     `json:"top25"`
	Matches         int     `json:"matches"`
	Top3            int     `json:"top3"`
	DateCollected   string  `json:"dateCollected"`
	Score           int     `json:"score"`
	Platform        int     `json:"platform"`
	TRNRating       float64 `json:"trnRating"`
	TRNRatingChange float64 `json:"trnRatingChange"`
}
