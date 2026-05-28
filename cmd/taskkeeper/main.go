package main

import (
	"log"
	"net/http"

	"github.com/dontagr/taskkeeper/internal/httpapi"
	"github.com/dontagr/taskkeeper/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	srv := httpapi.NewServer(store)

	addr := ":8080"
	log.Printf("taskkeeper listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Handler()); err != nil {
		log.Fatal(err)
	}
}
