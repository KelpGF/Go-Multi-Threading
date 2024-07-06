package services

import (
	"encoding/json"
	"net/http"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func SearchCepViaCep(cep string, channel chan<- interface{}) {
	var data ViaCepResponse

	// time.Sleep(1 * time.Second)
	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		channel <- data
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		channel <- data
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		channel <- data
		return
	}

	channel <- data
}

func PrintViaCepResponse(data ViaCepResponse) {
	println("CEP:", data.Cep)
	println("Logradouro:", data.Logradouro)
	println("Complemento:", data.Complemento)
	println("Unidade:", data.Unidade)
	println("Bairro:", data.Bairro)
	println("Localidade:", data.Localidade)
	println("UF:", data.Uf)
	println("IBGE:", data.Ibge)
	println("GIA:", data.Gia)
	println("DDD:", data.Ddd)
	println("SIAFI:", data.Siafi)
}
