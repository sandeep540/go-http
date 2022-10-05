package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Into Main")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Hello , %s!", r.URL.Path[1:])
}
