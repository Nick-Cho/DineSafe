package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Nick-Cho/allergy-project/internal/config"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{}

type AccInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

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
		log.Println("error unmarshalling request body from login request | ", err)
		response := events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: string("Error unmarshalling response body from login request"),
		}
		return response, nil
	}

	email := requestBody["email"]
	password := requestBody["password"]

	fmt.Printf("Request email : %s\n", email) //temp

	sqlRequest := fmt.Sprintf("SELECT * FROM allergy_db.Users WHERE Email='%s'", email)
	fmt.Printf("sql GET request: %s\n", sqlRequest)
	res, err := db.Query(sqlRequest)

	// Formatting MySQL response to JSON
	var acc AccInfo
	for res.Next() {
		// for each row, scan the result into our tag composite object
		var tempEmail string

		err = res.Scan(&acc.ID, &acc.Name, &tempEmail, &acc.Password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// log.Printf(acc.Name)
	}

	// fmt.Printf("Response from db execution: %s\n", res)

	if err != nil {
		log.Println("Error pulling user", err)
		return responses.ServerError(err), fmt.Errorf("error pulling user from table: %s", err)
	}
	defer db.Close()

	// Compare encrypted password pulled from DB
	byteStrDbPswd := []byte(acc.Password)

	err = bcrypt.CompareHashAndPassword(byteStrDbPswd, []byte(password))
	log.Println("Value received from bcrypt compare function: ", err)
	if err == nil {
		//Succesful login
		response := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: string("Success"),
		}
		return response, nil
	} else {
		response := events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: string("Failed login password does not match"),
		}
		return response, nil
	}

}
