package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	config "github.com/Nick-Cho/allergy-project/internal/config"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string

	// log.Println(request.Body)

	db := config.Connect()

	if request.Body == "" {
		log.Println("No request body provided")
		response := events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: string("No request body provided"),
		}
		return response, nil
	}

	sDec, _ := b64.StdEncoding.DecodeString(request.Body)

	err := json.Unmarshal([]byte(sDec), &requestBody)

	if err != nil {
		log.Println("error unmarshalling response body from insert restaurant request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from insert restaurant request")
	}

	address := requestBody["street_address"]
	name := requestBody["name"]
	fmt.Printf("Requested Insert Restaurant: %s, %s\n", address, name)

	sqlRequest := fmt.Sprintf("INSERT INTO allergy_db.Restaurants (street_address, name) VALUES ('%s', '%s')", address, name)
	fmt.Printf("sql POST request: %s\n", sqlRequest)
	res, err := db.Exec(sqlRequest)

	if err != nil {
		log.Println("error creating new restaurant", err)
		return responses.ServerError(err), fmt.Errorf("error inserting new entry into user table: %s", err)
	}

	lastId, err := res.LastInsertId()
	fmt.Printf("User ID inserted: %d\n", lastId)

	defer db.Close()

	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: "success",
	}

	return response, nil
}
