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
	// var requestBody map[string]string
	// request.Body =
	// 	`{
	// 		"latitude": "43.8272",
	// 		"longitude": "-79.2788992"
	// 	}`
	// err := json.Unmarshal([]byte(request.Body), &requestBody)

	// if err != nil {
	// 	log.Println("error unmarshalling lat and long | ", err)
	// 	return responses.ServerError(err), fmt.Errorf("error unmarshalling lat and long")
	// }

	longitude := request.QueryStringParameters["longitude"]
	latitude := request.QueryStringParameters["latitude"]

	log.Println("Latitude from query string: ", latitude)
	log.Println("Longitude from query string: ", longitude)

	if latitude == "" || longitude == "" {
		response := events.APIGatewayProxyResponse{
			StatusCode: 202,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     "*",
				"Access-Control-Allow-Credentials": "true",
			},
			Body: "latitude and/or Longitude not provided to getNearbyRestaurants route",
		}
		return response, nil
	}

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=10000&type=restaurant&key=%s", latitude, longitude, os.Getenv("GOOGLE_API_KEY"))

	// log.Println("URL without variables: ", presetUrl)
	log.Println("URL with variables: ", url)

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
