package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dmaxim/hello-api/handlers"
	"github.com/dmaxim/hello-api/handlers/rest"
	"github.com/dmaxim/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}
	mux := http.NewServeMux()

	tranlationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(tranlationService)

	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
