package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "seu-projeto/internal/services"
    "seu-projeto/internal/models"
    "go.opentelemetry.io/otel"
)

func HandleTemperature(w http.ResponseWriter, r *http.Request) {
    ctx, span := otel.Tracer("service-b").Start(r.Context(), "handle-temperature")
    defer span.End()

    vars := mux.Vars(r)
    cep := vars["cep"]

    // Validação do CEP
    if !isValidCep(cep) {
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid zipcode"})
        return
    }

    // Busca informações do CEP
    cepInfo, err := services.GetCepInfo(ctx, cep)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "can not find zipcode"})
        return
    }

    // Busca temperatura
    temp, err := services.GetWeather(ctx, cepInfo.City)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "error fetching weather"})
        return
    }

    response := models.TemperatureResponse{
        City:   cepInfo.City,
        TempC:  temp.Celsius,
        TempF:  temp.Celsius*1.8 + 32,
        TempK:  temp.Celsius + 273.15,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
} 