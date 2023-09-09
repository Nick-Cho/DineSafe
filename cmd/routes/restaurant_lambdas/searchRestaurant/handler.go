package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	search := request.QueryStringParameters["search"]
	longitude := request.QueryStringParameters["longitude"]
	latitude := request.QueryStringParameters["latitude"]

	log.Println("search parameter: ", search)
	log.Println("longitude parameter: ", longitude)
	log.Println("latitude parameter: ", latitude)

	if search == "" || longitude == "" || latitude == "" {
		response := events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: "Query string values missing in searchRestaurant route. Must provide a valid search longitude and a latitude",
		}
		return response, nil
	}
	// url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%s&inputtype=textquery&fields=formatted_address,name,rating,opening_hours&key=%s", search, os.Getenv("GOOGLE_API_KEY"))
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?fields=formatted_address,name,rating,opening_hours&input=%s&inputtype=textquery&locationbias=circle:20000@%s,%s&key=%s", search, latitude, longitude, os.Getenv("GOOGLE_API_KEY"))
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
