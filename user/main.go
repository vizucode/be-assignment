package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello From User Service"))
	})

	fmt.Println("Server running on port 8082")
	http.ListenAndServe(":8082", nil)
}
