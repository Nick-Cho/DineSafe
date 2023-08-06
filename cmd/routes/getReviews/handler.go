package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Nick-Cho/allergy-project/internal/config"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

type RestaurantInfo struct {
	Street_Address int    `json:"id"`
	Name           string `json:"name"`
	City           string `json:"city"`
}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	request.Body =
		`{
			"street_address": "5285 Yonge St Unit5"
		}`
	db := config.Connect()

	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	streetAddress := requestBody["street_address"]
	fmt.Printf("Request street address : %s\n", streetAddress) //temp

	sqlRequest := fmt.Sprintf("SELECT * FROM allergy_db.Restaurant WHERE street_address='%s'", streetAddress)
	fmt.Printf("sql GET request: %s\n", sqlRequest)
	res, err := db.Query(sqlRequest)

	// Formatting MySQL response to JSON
	var restaurant RestaurantInfo
	for res.Next() {
		// for each row, scan the result into our tag composite object
		var tempStreetAddress string
		var tempName string
		var tempCity string
		err = res.Scan(&restaurant.Street_Address, &restaurant.Name, &restaurant.City, &tempStreetAddress, &tempName, &tempCity)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		log.Printf(restaurant.Name)
	}

	// fmt.Printf("Response from db execution: %s\n", res)

	if err != nil {
		log.Println("Error pulling user", err)
		return responses.ServerError(err), fmt.Errorf("error pulling user from table: %s", err)
	}
	defer db.Close()

	responseBody, err := json.Marshal(restaurant)
	if err != nil {
		log.Println("ERROR MARSHALLING RESPONSE BODY TO JSON", err)
		return responses.ServerError(err), fmt.Errorf("ERROR MARSHALLING RESPONSE BODY TO JSON")
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseBody),
	}

	return response, nil
}
