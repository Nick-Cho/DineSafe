package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Nick-Cho/allergy-project/internal/config"
	userModel "github.com/Nick-Cho/allergy-project/internal/models"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// fmt.Println("Hello World")
	var requestBody map[string]string
	var user userModel.User
	fmt.Println(user)
	db := config.Connect()
	defer db.Close()

	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		// return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from spotify update user request")
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
	}

	return response, nil
}
