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

	log.Printf("Server is running on port %s", args[1])
	serverStr := fmt.Sprintf(":%s", args[1])
	http.ListenAndServe(serverStr, nil)
}
