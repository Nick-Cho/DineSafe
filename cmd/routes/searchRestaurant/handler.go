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
	// 		"search": "Yunshang",
	// 		"latitude": "43.8272",
	// 		"longitude": "-79.2788992"
	// 	}`
	// err := json.Unmarshal([]byte(request.Body), &requestBody)

	// if err != nil {
	// 	log.Println("error unmarshalling lat and long | ", err)
	// 	return responses.ServerError(err), fmt.Errorf("error unmarshalling lat and long")
	// }

	search := request.QueryStringParameters["search"]
	longitude := request.QueryStringParameters["longitude"]
	latitude := request.QueryStringParameters["latitude"]
	log.Println("search parameter: ", search)
	log.Println("longitude parameter: ", longitude)
	log.Println("latitude parameter: ", latitude)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%s&inputtype=textquery&locationbias=circle:20000@%s,%s&fields=formatted_address,name,rating,opening_hours&key=%s", search, latitude, longitude, os.Getenv("GOOGLE_API_KEY"))
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
