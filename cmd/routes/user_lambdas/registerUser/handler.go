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
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{}

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string

	if request.Body == "" {
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

	db := config.Connect()

	sDec, _ := b64.StdEncoding.DecodeString(request.Body)
	log.Println("login request body: ", sDec)
	err := json.Unmarshal([]byte(sDec), &requestBody)

	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	//Should check if the user exists first
	email := requestBody["email"]
	name := requestBody["name"]
	password := []byte(requestBody["password"])

	fmt.Printf("Request email name and password: %s, %s, %s\n", email, name, password) //temp

	//Encrypt password before saving it in DB
	hashedPswd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		log.Println("error encrypting password | ", err)
		return responses.ServerError(err), fmt.Errorf("error encrypting password")
	}

	convertedHashedPswd := string(hashedPswd)

	sqlRequest := fmt.Sprintf("INSERT INTO allergy_db.Users (id, name, email, password) VALUES (2, '%s', '%s', '%s')", name, email, convertedHashedPswd)
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
