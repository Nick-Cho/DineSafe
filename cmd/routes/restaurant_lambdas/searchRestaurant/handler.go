package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

	// Replacing white space with ASCII encoded space (%20)
	replacementChar := "%20"
	replacedStr := strings.ReplaceAll(search, " ", string(replacementChar))

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
	// ASCII encoded characters
	comma := "%2C"
	colon := "%3A"
	address_sign := "%40"

	proper_url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?fields=formatted_address%2Cname%2Crating%2Copening_hours&input=yunshang&inputtype=textquery&locationbias=circle%3A20000%4043.4730755%2C-80.5395694&key=AIzaSyCEMyZMx4vfrx8-fU22fwGljlPOBkEervo"
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?fields=formatted_address%sname%srating%sopening_hours&input=%s&inputtype=textquery&locationbias=circle%s20000%s%s%s%s&key=%s", comma, comma, comma, replacedStr, colon, address_sign, latitude, comma, longitude, os.Getenv("GOOGLE_API_KEY"))
	log.Println("proper url call from api test website: ", proper_url)
	log.Println("url call without ASCII characters: ", url)

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
