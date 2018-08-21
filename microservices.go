package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eoinahern/new-cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	fmt.Println("running .....")
	http.ListenAndServe(":8080", nil)

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return "8080"
	}

	return ":" + port
}

func echo(w http.ResponseWriter, req *http.Request) {

	message := req.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "message is = "+message)

}

func index(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello cloud native go!!!")
}
