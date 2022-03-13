package handler

import (
	"encoding/json"
	"net/http"
	"tesla-take-home-challenge/internal/entity"
)

type AgeRequest struct {
	Age int `json:"age"`
}

// age handler
func HandlerAge(i *entity.Inventory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// post
		if r.Method == "POST" {
			// get request body
			var ageRequest AgeRequest
			err := json.NewDecoder(r.Body).Decode(&ageRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// set value
			i.X = ageRequest.Age
			var reps = map[string]bool{
				"ok": true,
			}
			// response
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(&reps)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// get
		if r.Method == "GET" {
			var reps = map[string]int{
				"age": i.X,
			}
			// response
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(&reps)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		}

	}
}

// car handler
func HandlerCar(i *entity.Inventory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// get
		if r.Method == "GET" {
			car, err := i.SellCar()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var reps = map[string]int64{
				"car": car,
			}
			// response
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(&reps)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		}

	}
}

// rate handler
func HandlerRate(i *entity.Inventory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get
		if r.Method == "GET" {

			var reps = map[string]int{
				"rate": i.R,
			}
			// response
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(&reps)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		}

	}
}

// buffer handler
func HandlerBuffer(i *entity.Inventory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get
		if r.Method == "GET" {

			var reps = map[string]int{
				"buffer": i.GetN(),
			}
			// response
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(&reps)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		}

	}
}
