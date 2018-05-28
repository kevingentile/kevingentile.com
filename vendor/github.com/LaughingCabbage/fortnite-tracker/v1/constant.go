package tracker

import (
	"errors"
	"time"
)

// TestResourcePath holds the path to test resources
const testResourcePath = "test-resource/"

// TestKey is a LIVE KEY. Should limit requests
const testKey = ".tracker-key"

// requestLimitDuration is the amout of time required between requests. 0.5 per second
const requestLimitDuration = time.Second * 2

// headerKeyName is the header fild name used to supply an API token to fortnite tracker
const headerKeyName = "TRN-Api-Key"

// GET https://api.fortnitetracker.com/v1/profile/{platform}/{epic-nickname}

//Version contains the current API version for fornite tracker
const Version = "v1"

//ProfileRoute contains the core route for requesting profile resources
const ProfileRoute = "https://api.fortnitetracker.com/v1/profile/"

// ErrNotFound is an error type used when entries are not found
var ErrNotFound = errors.New("entry not found")
