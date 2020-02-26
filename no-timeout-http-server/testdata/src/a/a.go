package a

import (
	"net/http"
	"time"
)

func main() {
	_ = http.Server{} // want "Do not use net/http.Server with no ReadTimeout" "Do not use net/http.Server with no WriteTimeout"
	_ = http.Server{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	_ = &http.Server{} // want "Do not use net/http.Server with no ReadTimeout" "Do not use net/http.Server with no WriteTimeout"
}
