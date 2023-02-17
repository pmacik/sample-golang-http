package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", HelloServer)
	fmt.Printf("Listening on port %d... \n", 8080)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Hello, %s!", r.URL.Path[7:])
	fmt.Println(message)
	fmt.Fprintf(w, "%s", message)
}
