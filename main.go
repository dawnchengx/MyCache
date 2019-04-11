package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MyCache/objects"
)

func main() {
	fmt.Println()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
