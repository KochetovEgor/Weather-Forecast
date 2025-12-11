package app

import (
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/index.html")
}

func Run() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/forecast", forecast)
	http.ListenAndServe("localhost:8000", nil)
}
