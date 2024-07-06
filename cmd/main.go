package main

import (
	"os"
	"time"

	"github.com/KelpGF/Go-Multi-Threading/internal/services"
)

func main() {
	zipCode := os.Args[1]
	if zipCode == "" {
		panic("Zip code is required")
	}

	channelBrAPI := make(chan interface{})
	go services.SearchCepBrApi(zipCode, channelBrAPI)

	channelViaCep := make(chan interface{})
	go services.SearchCepViaCep(zipCode, channelViaCep)

	select {
	case response := <-channelBrAPI:
		println(">>> BrAPI \n")
		data := response.(services.BrApiResponse)
		services.PrintBrApiResponse(data)
	case response := <-channelViaCep:
		println(">>> ViaCep \n")
		data := response.(services.ViaCepResponse)
		services.PrintViaCepResponse(data)
	case <-time.After(1 * time.Second):
		println("Timeout")
	}
}
