package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello From Payment Service"))
	})

	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8081", nil)
}
