package chart

import {
	"fmt"
	"net/http"
	"string"
}

func GetChartData(response http.ResponseWriter, request *http.Request) {
	type, err := request.URL.Query()["type"]
	if err == nil {
		return
	}
	var dataFile string = string(type)
	fmt.Println("dataFile:", dataFile)
	fmt.Println("rootDir:", "static/" + dataFile + ".json")

	fs := http.FileServer(http.Dir("static/" + dataFile + ".json"))

	return fs()
}
