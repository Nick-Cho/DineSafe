package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := godotenv.Load(".env")
	log.Printf("Error from loading env: %s\n", err)

	log.Printf("API KEY: %s", os.Getenv("GOOGLE_API_KEY"))

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Restaurant&inputtype=textquery&fields=formatted_address,name,rating,opening_hours,geometry&key=%s", os.Getenv("GOOGLE_API_KEY"))
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// var requestBody map[string]string
	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(body),
	}
	return response, nil
}
