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

	http.ServeFile(response, request, "static/"+dataFile)
}
