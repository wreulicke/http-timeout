package a

import (
	"net/http"
	"time"
)

func main() {
	_ = http.Client{}  // want "Do not use net/http.Client with no Timeout"
	_ = &http.Client{} // want "Do not use net/http.Client with no Timeout"
	_ = http.Client{
		Timeout: 20 * time.Second,
	}
}
