package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hoangluan97/echart/backend/pkg/chart"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/static", chart.GetChartData).Methods("GET")
	spa := chart.SpaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	return router
}

func main() {
	router := initRouter()

	router.Use(mux.CORSMethodMiddleware(router))

	fmt.Println("connected")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
