package services

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "seu-projeto/internal/models"
    "go.opentelemetry.io/otel"
)

func GetCepInfo(ctx context.Context, cep string) (*models.CepInfo, error) {
    _, span := otel.Tracer("service-b").Start(ctx, "get-cep-info")
    defer span.End()

    url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("CEP não encontrado")
    }

    var cepInfo models.CepInfo
    if err := json.NewDecoder(resp.Body).Decode(&cepInfo); err != nil {
        return nil, err
    }

    return &cepInfo, nil
} 