package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	request.Body =
		`{
			"latitude": "43.8272",
			"longitude": "-79.2788992",
		}`
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	// latitude := requestBody["latitude"]
	// longitude := requestBody["longitude"]
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=43.8272,-79.2788992&radius=10000&type=restaurant&key=AIzaSyCEMyZMx4vfrx8-fU22fwGljlPOBkEervo")
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
