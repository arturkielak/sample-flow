package main
import (
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Przekopiuj całe body requestu do odpowiedzi
	io.Copy(w, r.Body)
}

func main() {
	http.HandleFunc("/", echoHandler)

	// Start serwera HTTP na porcie 8080
	http.ListenAndServe(":8080", nil)
}