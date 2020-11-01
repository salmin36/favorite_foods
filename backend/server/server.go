package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type Server struct {
	databaseAddress  string
	databasePassword string
	rdb              *redis.Client
}

type Foods struct {
	Name        string
	Ingridients []string
	Steps       []string
	Comments    []string
}

func NewServer(rdb *redis.Client) *Server {
	server := Server{
		databaseAddress:  "redis:6379",
		databasePassword: "temp_pass_123",
		rdb:              rdb,
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
	var ctx = context.Background()
	val, err := server.rdb.LLen(ctx, "foods").Result()
	if err != nil {
		io.WriteString(w, "{\"Error\":\""+err.Error()+"\"}")
		return
	} else {
		//io.WriteString(w, "{\"Name of list\": \"foods\", \"Size\": \""+strconv.FormatInt(val, 10)+"\"}")
		fmt.Println("Size of list: " + strconv.FormatInt(val, 10))
	}

	//val, err := server.rdb.Get(ctx, "foods").Result()
	results, err := server.rdb.LRange(ctx, "foods", 0, val).Result()
	if err != nil {
		fmt.Println(err.Error())
		io.WriteString(w, "{\"Error\":\"Could not find any foods\"}")
	} else {
		fmt.Println("Got something for redis")
		io.WriteString(w, results[0])
	}
	//foods := server.fakeDataBaseFood()
	//json.NewEncoder(w).Encode(foods)
	//"{\"Name\":\"Kinkkukiusaus\", \"Ingridients\":[\"Kinkkusuikaleita 200g\",\"Kermaa 4dl\",\"Pippuria \",\"Suolaa\",\"Sipuliperunasekoitus 300g\"],\"Steps\":[\"Sekoita kinkkusuikaleet ja sipuliperunasekoitukset sekä kaikki muut samaan vuokaan.\",\"Paista 200 asteessa 30-35 min\"],\"Comments\":[\"Hyvää ruokaa!\",\"Toimiva\"]}"
}

func (server *Server) fakeDataBaseFood() *Foods {
	foods := Foods{
		Name:        "Kinkkukiusaus",
		Ingridients: []string{"Kinkkusuikaleita 200g", "Kermaa 4dl", "Pippuria ", "Suolaa", "Sipuliperunasekoitus 300g"},
		Steps:       []string{"Sekoita kinkkusuikaleet ja sipuliperunasekoitukset sekä kaikki muut samaan vuokaan.", "Paista 200 asteessa 30-35 min"},
		Comments:    []string{"Hyvää ruokaa!", "Toimiva"},
	}
	return &foods
}

func (server *Server) AddFoodEndpoint(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) ModifyFoodEndpoint(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) RemoveFoodEndpoint(w http.ResponseWriter, r *http.Request) {

}
