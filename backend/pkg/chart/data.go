package chart

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type SpaHandler struct {
	staticPath string
	indexPath  string
}

func GetChartData(response http.ResponseWriter, request *http.Request) {
	chartType := request.URL.Query().Get("type")
	dataFile := chartType + ".json"
	fmt.Println("rootDir:", "/static/"+dataFile)
	allowOrigin(response)
	http.ServeFile(response, request, "static/"+dataFile)
}

func (h SpaHandler) RenderHomePage(response http.ResponseWriter, request *http.Request) {
	path, err := filepath.Abs(request.URL.Path)
	fmt.Println("path:", path)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)
	fmt.Println("path pretend:", path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(response, request, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(response, request)
}

func allowOrigin(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Methods", http.MethodGet)
}
