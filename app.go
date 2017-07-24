package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	fmt.Fprintf(w, "Hello, from: %s\n", name)
}
func main() {
	http.HandleFunc("/", indexHandle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
