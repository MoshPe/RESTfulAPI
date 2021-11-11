package main

import (
	cli "clientAPI.com/client/go-client/client"
	"context"
	"log"
	"net/http"
	"time"
)

func main(){
	ctx := context.Background()
	config := cli.Configuration{
		BasePath:      "http://localhost:8080/v2",
		Host:          "client",
		Scheme:        "",
		DefaultHeader: nil,
		UserAgent:     "",
		HTTPClient:    &http.Client{Timeout: time.Second* 10},
	}
	client := cli.NewAPIClient(&config)
	newPet := cli.Pet{
		Id:        318,
		Category:  nil,
		Name:      "Snoopi",
		PhotoUrls: nil,
		Tags:      nil,
		Status:    "available",
	}
	response, err := client.PetApi.AddPet(ctx,newPet)
	if err != nil {
		log.Fatal("Error in Adding a new pet")
	}
	log.Printf(response.Status)
}
