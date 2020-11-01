package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	databaseAddress  string
	databasePassword string
}

type Foods struct {
	Name        string
	Ingridients []string
	Steps       []string
	Comments    []string
}

func NewServer() *Server {
	server := Server{
		databaseAddress:  "redis:6379",
		databasePassword: "temp_pass_123",
	}

	return &server
}

func (server *Server) StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/foods", server.FoodsEndpoint).Methods("GET")
	http.ListenAndServe(":8081", router)
}

func (server *Server) FoodsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foods endpoint called")
	foods := Foods{
		Name:        "Kinkkukiusaus",
		Ingridients: []string{"Kinkkusuikaleita 200g", "Kermaa 4dl", "Pippuria ", "Suolaa", "Sipuliperunasekoitus 300g"},
		Steps:       []string{"Sekoita kinkkusuikaleet ja sipuliperunasekoitukset sekä kaikki muut samaan vuokaan.", "Paista 200 asteessa 30-35 min"},
		Comments:    []string{"Hyvää ruokaa!", "Toimiva"},
	}

	json.NewEncoder(w).Encode(foods)
}
