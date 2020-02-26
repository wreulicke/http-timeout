package a

import (
	"net/http"
)

func main() {
	f := func(http.ResponseWriter, *http.Request) {
	}
	http.Handle("/", http.HandlerFunc(f))                           // want "Do not use net/http.Handle because default http server has no timeout"
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) { // want "Do not use net/http.HandleFunc because default http server has no timeout"
	})
	http.ListenAndServe(":8080", nil) // want "Do not use net/http.ListenAndServe because default http server has no timeout"
}
