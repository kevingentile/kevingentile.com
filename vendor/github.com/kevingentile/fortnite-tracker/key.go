// Copyright 2019 Kevin Gentile.
// Licensed under GNU General Public License v3.0

package tracker

import "io/ioutil"

// Key is your API token
type Key struct {
	Value string
}

//LoadKey is used to load an API token from a file.
func LoadKey(path string) *Key {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return &Key{Value: string(data)}
}
