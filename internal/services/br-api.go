package services

import (
	"encoding/json"
	"net/http"
)

type BrApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func SearchCepBrApi(cep string, channel chan<- interface{}) {
	var data BrApiResponse

	// time.Sleep(1 * time.Second)
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
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

func PrintBrApiResponse(data BrApiResponse) {
	println("CEP:", data.Cep)
	println("State:", data.State)
	println("City:", data.City)
	println("Neighborhood:", data.Neighborhood)
	println("Street:", data.Street)
	println("Service:", data.Service)
}
