package entity

type BrasilApiCep struct {
	Cep        string `json:"cep"`
	Estado     string `json:"state"`
	Cidade     string `json:"city"`
	Bairro     string `json:"neighborhood"`
	Logradouro string `json:"street"`
	Servico    string `json:"service"`
}

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
