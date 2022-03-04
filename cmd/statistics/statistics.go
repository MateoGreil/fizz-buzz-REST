package statistics

import (
	"encoding/json"
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
		FizzBuzz: fizzBuzz,
		Hits:     2, //TODO
	}
	json.NewEncoder(w).Encode(response)
}
