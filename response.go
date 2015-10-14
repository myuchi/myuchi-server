package main

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, resp interface{}) error {
	var data []byte
	var err error
	if data, err = json.Marshal(resp); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
	return nil
}
