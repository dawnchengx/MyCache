package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dawnchengx/MyCache/heartbeat"
	"github.com/dawnchengx/MyCache/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
