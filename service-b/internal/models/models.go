package models

type CepInfo struct {
    Cep         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Complemento string `json:"complemento"`
    Bairro      string `json:"bairro"`
    City        string `json:"localidade"`
    Uf          string `json:"uf"`
}

type Temperature struct {
    Celsius float64
}

type TemperatureResponse struct {
    City  string  `json:"city"`
    TempC float64 `json:"temp_C"`
    TempF float64 `json:"temp_F"`
    TempK float64 `json:"temp_K"`
} 