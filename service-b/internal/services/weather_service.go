package services

import (
	"cep-opentelemetry/service-b/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
)

func GetWeather(ctx context.Context, city string) (*models.Temperature, error) {
    _, span := otel.Tracer("service-b").Start(ctx, "get-weather")
    defer span.End()

    apiKey := os.Getenv("WEATHER_API_KEY")
    url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var weatherResp struct {
        Current struct {
            TempC float64 `json:"temp_c"`
        } `json:"current"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
        return nil, err
    }

    return &models.Temperature{
        Celsius: weatherResp.Current.TempC,
    }, nil
} 