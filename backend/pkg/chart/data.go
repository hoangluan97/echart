package chart

import (
	"fmt"
	"net/http"
)

func GetChartData(response http.ResponseWriter, request *http.Request) {
	chartType := request.URL.Query().Get("type")
	dataFile := chartType + ".json"
	fmt.Println("dataFile:", dataFile)
	fmt.Println("rootDir:", "/static/"+dataFile+".json")
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	http.ServeFile(response, request, "static/"+dataFile)
}
