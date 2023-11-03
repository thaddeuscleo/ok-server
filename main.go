package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request Accepted")
		jsonByte, _ := json.Marshal("ok")
		w.Write(jsonByte)
	})

	port := "8080"
	if len(args) > 1 {
		port = args[1]
	}

	log.Printf("Server is running on port %s", port)
	serverStr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(serverStr, nil)
}
