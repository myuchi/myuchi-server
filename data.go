package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadData(path string, data interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func saveData(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	return err
}
