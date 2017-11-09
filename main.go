package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)
	log.Printf("Listening on %s...\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}
