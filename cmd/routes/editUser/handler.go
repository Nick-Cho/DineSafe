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

type AccInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	// request.Body =
	// 	`{
	// 		"email": "nicholas.cho@hotmail.ca"
	// 	}`
	db := config.Connect()

	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	email := requestBody["email"]
	fmt.Printf("Request email : %s\n", email) //temp

	sqlRequest := fmt.Sprintf("SELECT * FROM allergy_db.Users WHERE Email='%s'", email)
	fmt.Printf("sql GET request: %s\n", sqlRequest)
	res, err := db.Query(sqlRequest)

	// Formatting MySQL response to JSON
	var acc AccInfo
	for res.Next() {
		// for each row, scan the result into our tag composite object
		var tempEmail string
		var tempPassword string
		err = res.Scan(&acc.ID, &acc.Name, &tempEmail, &tempPassword)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		log.Printf(acc.Name)
	}

	// fmt.Printf("Response from db execution: %s\n", res)

	if err != nil {
		log.Println("Error pulling user", err)
		return responses.ServerError(err), fmt.Errorf("error pulling user from table: %s", err)
	}
	defer db.Close()

	responseBody, err := json.Marshal(acc)
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
