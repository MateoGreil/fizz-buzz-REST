package main

import (
	"log"
	"mateogreil/fizz-buzz-REST/cmd/config"
	"mateogreil/fizz-buzz-REST/cmd/download"
	"mateogreil/fizz-buzz-REST/cmd/fizz_buzz"
	"mateogreil/fizz-buzz-REST/cmd/models"
	"mateogreil/fizz-buzz-REST/cmd/statistics"
	"net/http"
)

func autoMigrateModels() {
	config.Db().AutoMigrate(&models.FizzBuzz{})
	config.Db().AutoMigrate(&models.Query{})
}

func main() {
	config.DatabaseInit()
	autoMigrateModels()
	http.HandleFunc("/fizz-buzz", fizz_buzz.Handler)
	http.HandleFunc("/fizz-buzz/statistics", statistics.Handler)
	http.HandleFunc("/fizz-buzz/download", download.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
