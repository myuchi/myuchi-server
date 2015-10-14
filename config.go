package main

import "path/filepath"

var config struct {
	Addr         string `json:"addr"`
	RootPassword string `json:"root_password"`
}

func loadConfig() error {
	return loadData(filepath.Join(database, "config.json"), &config)
}
