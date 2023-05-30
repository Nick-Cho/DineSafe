package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Nick-Cho/allergy-project/internal/config"
	userModel "github.com/Nick-Cho/allergy-project/internal/models"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	var user userModel.User

	fmt.Println(user) //temp

	db := config.Connect()
	defer db.Close()

	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	//Should check if the user exists first
	email := requestBody["email"]
	name := requestBody["name"]
	password := requestBody["password"]

	//Encrypt password before saving it in DB

	_, err = db.Exec("INSERT INTO user(name, email, password) VALUES(?, ?, ?)", name, email, password)
	fmt.Println(email, name, password) //temp

	if err != nil {
		log.Println("error creating new user", err)
		return responses.ServerError(err), fmt.Errorf("error inserting new entry into user table")
	}
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
