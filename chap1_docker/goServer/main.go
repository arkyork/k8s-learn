package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	// ハンドラを登録
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker!")
	})

	log.Println("Starting server on :8080")
	
	// サーバを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}