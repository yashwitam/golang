package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/ymishra/BookService/api"
)

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.CreateListAllBooksHandleFunc)
	http.HandleFunc("/api/books/", api.UpdateDeleteBookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string{
	os.Clearenv()
	port := os.Getenv("PORT")
	if(len(port) == 0){
		port = "8080"
	}
	return ":" + port;
}
func index(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from cloud native go!")
}

func echo(w http.ResponseWriter, r *http.Request){
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text-plain")
	fmt.Fprint(w, message)
}