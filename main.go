package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Car struct (Model)
type Car struct {
	ID     string  `json:"id"`
	Make   string  `json:"make"`
	Model  string  `json:"model"`
	Type *CarType `json:"type"`
}

// CarType struct
type CarType struct {
	Typeofcar string `json:"Typeofcar"`
	Numberofdoors  int `json:"numberofdoors"`
}

// Init cars var as a slice car struct
var cars []Car

// Get all cars
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// Get single car
func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through cars and find one with the id from the params
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

// Add new car
func createcar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

// Update car
func updatecar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return
		}
	}
}

// Delete car
func deletecar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	cars = append(cars, car{ID: "1", Make: "438227", Model: "car One", Cartype: &Cartype{Typeofcar: "John", Numberofdoors: "Doe"}})
	cars = append(cars, car{ID: "2", Make: "454555", Model: "car Two", Cartype: &Cartype{Typeofcar: "Steve", Numberofdoors: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/cars", getcars).Methods("GET")
	r.HandleFunc("/cars/{id}", getcar).Methods("GET")
	r.HandleFunc("/cars", createcar).Methods("POST")
	r.HandleFunc("/cars/{id}", updatecar).Methods("PUT")
	r.HandleFunc("/cars/{id}", deletecar).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"Make":"Subaru",
// 	"Model":"Legacy",
// 	"Cartype":{"Typeofcar":"Sedan","Numberofdoors":4}
// }
