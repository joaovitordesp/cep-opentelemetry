package main

import (
    "context"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "seu-projeto/internal/handlers"
    "seu-projeto/pkg/telemetry"
)

func main() {
    // Inicializa o tracer
    tp, err := telemetry.InitTracer()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Printf("Error shutting down tracer provider: %v", err)
        }
    }()

    r := mux.NewRouter()
    r.HandleFunc("/temperature/{cep}", handlers.HandleTemperature).Methods("GET")
    
    log.Println("Service B starting on port 8081")
    log.Fatal(http.ListenAndServe(":8081", r))
} 