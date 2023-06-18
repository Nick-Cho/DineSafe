package main

import (
	"encoding/json"
	"fmt"
	"log"

	config "github.com/Nick-Cho/allergy-project/internal/config"
	userModel "github.com/Nick-Cho/allergy-project/internal/models"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	var user userModel.User
	request.Body =
		`{
			"name": "Nick",
			"email": "nicholas.cho@hotmail.ca",
			"password": "123"
		}`

	log.Println(request.Body)
	fmt.Println(user) //temp

	db := config.Connect()

	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	//Should check if the user exists first
	email := requestBody["email"]
	name := requestBody["name"]
	password := requestBody["password"]
	fmt.Printf("Request email name and password: %s, %s, %s\n", email, name, password) //temp

	//Encrypt password before saving it in DB

	sqlRequest := fmt.Sprintf("INSERT INTO allergy_db.Users (id, name, email, password) VALUES (2, '%s', '%s', '%s')", name, email, password)
	fmt.Printf("sql POST request: %s\n", sqlRequest)
	res, err := db.Exec(sqlRequest)

	if err != nil {
		log.Println("error creating new user", err)
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
