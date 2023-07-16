package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	request.Body =
		`{
			"latitude": "43.8272",
			"longitude": "-79.2788992"
		}`
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		log.Println("error unmarshalling lat and long | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling lat and long")
	}

	latitude := requestBody["latitude"]
	longitude := requestBody["longitude"]

	// presetUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=43.8272,-79.2788992&radius=10000&type=restaurant&key=AIzaSyCEMyZMx4vfrx8-fU22fwGljlPOBkEervo")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=10000&type=restaurant&key=AIzaSyCEMyZMx4vfrx8-fU22fwGljlPOBkEervo", latitude, longitude)

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
