package statistics

import (
	"encoding/json"
	"log"
	"mateogreil/fizz-buzz-REST/cmd/models"
	"net/http"
)

type Response struct {
	FizzBuzz     models.FizzBuzz
	QueriesCount int
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var fizzBuzz = models.FizzBuzz{}
	models.MaxRequestedFizzBuzz(&fizzBuzz)
	response := Response{
		FizzBuzz:     fizzBuzz,
		QueriesCount: 2,
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(responseJson)
}
