package a

import (
	"net/http"
)

func main() {
	http.Get("test")           // want "Do not use net/http.Get because default client has no timeout"
	http.Post("test", "", nil) // want "Do not use net/http.Post because default client has no timeout"
	http.PostForm("test", nil) // want "Do not use net/http.PostForm because default client has no timeout"
	_ = http.DefaultClient     // want "Do not use net/http.DefaultClient because default client has no timeout"
}
