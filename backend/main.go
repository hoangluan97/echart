package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hoangluan97/echart/backend/pkg/chart"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/static/", chart.GetChartData)

	return router
}

func main() {
	router := initRouter()

	err := http.ListenAndServe(":8080", router)
	if err == nil {
		fmt.Println("connected")
	}
	if err != nil {
		panic(err)
	}
}
