package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// fmt.Println("Hello World")
	var requestBody map[string]string
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		// return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from spotify update user request")
	}
	db, err := sql.Open("mysql", "root:$Bigley2209@tcp(localhost:3306)/allergy_db")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db)
	defer db.Close()

	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
	}

	return response, nil
}
