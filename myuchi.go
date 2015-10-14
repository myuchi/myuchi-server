package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	database string
)

func main() {
	database = *flag.String("database", "./database", "database directory")
	flag.Parse()
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", statusHandler)
	http.HandleFunc("/user", userHandler)
	log.Println("Start Server")
	if err := http.ListenAndServe(config.Addr, nil); err != nil {
		log.Fatal(err)
	}
}
