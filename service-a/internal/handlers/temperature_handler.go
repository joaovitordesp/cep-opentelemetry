package handlers

import (
    "encoding/json"
    "net/http"
    "regexp"
    "seu-projeto/internal/models"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/propagation"
)

func HandleTemperatureRequest(w http.ResponseWriter, r *http.Request) {
    ctx, span := otel.Tracer("service-a").Start(r.Context(), "handle-temperature-request")
    defer span.End()

    var req models.CepRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    // Validação do CEP
    if !isValidCep(req.Cep) {
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid zipcode"})
        return
    }

    // Chamada para o Serviço B
    resp, err := forwardToServiceB(ctx, req.Cep)
    if err != nil {
        http.Error(w, "service unavailable", http.StatusServiceUnavailable)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(resp.StatusCode)
    w.Write(resp.Body)
}

func isValidCep(cep string) bool {
    match, _ := regexp.MatchString(`^\d{8}$`, cep)
    return match
} 