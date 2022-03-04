package fizz_buzz

import (
	"encoding/json"
	"mateogreil/fizz-buzz-REST/cmd/config"
	"mateogreil/fizz-buzz-REST/cmd/models"
	"net/http"
	"strconv"
)

// Main function of fizz buzz : it convert the query into the appropriate string
func response(fizzBuzz models.FizzBuzz) []string {
	var response []string
	for i := 1; i <= fizzBuzz.Limit; i++ {
		if i%fizzBuzz.Int1 == 0 {
			response = append(response, fizzBuzz.Str1)
		}
		if i%fizzBuzz.Int2 == 0 {
			response = append(response, fizzBuzz.Str2)
		}
		if i%fizzBuzz.Int1 != 0 && i%fizzBuzz.Int2 != 0 {
			response = append(response, strconv.Itoa(i))
		}
	}

	return response
}

// It handle the query and response the string
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	int1, e := strconv.Atoi(r.URL.Query().Get("int1"))
	if e != nil {
		int1 = 3
	}
	int2, e := strconv.Atoi(r.URL.Query().Get("int2"))
	if e != nil {
		int2 = 5
	}
	limit, e := strconv.Atoi(r.URL.Query().Get("limit"))
	if e != nil {
		limit = 100
	}
	str1 := r.URL.Query().Get("str1")
	if len(str1) == 0 {
		str1 = "fizz"
	}
	str2 := r.URL.Query().Get("str2")
	if len(str2) == 0 {
		str2 = "buzz"
	}

	fizzBuzz := models.FizzBuzz{Int1: int1, Int2: int2, Limit: limit, Str1: str1, Str2: str2}
	models.FindOrCreateFizzBuzz(&fizzBuzz)

	query := models.Query{Path: r.URL.Path, RemoteAddr: r.RemoteAddr, FizzBuzzId: fizzBuzz.ID}
	config.Db().Create(&query)
	response := response(fizzBuzz)
	json.NewEncoder(w).Encode(response)
}
