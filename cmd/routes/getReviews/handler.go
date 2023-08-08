package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Nick-Cho/allergy-project/internal/config"
	"github.com/Nick-Cho/allergy-project/internal/responses"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

type ReviewInfo struct {
	Id     int    `json:"id"`
	Review string `json:"review"`
}

type RestaurantInfo struct {
	StreetAddress string `json:"streetAddress"`
	Name          string `json:"name"`
	City          string `json:"city"`
}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
	request.Body =
		`{
			"street_address": ""
		}`
	db := config.Connect()

	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		log.Println("error unmarshalling response body from register user request | ", err)
		return responses.ServerError(err), fmt.Errorf("error unmarshalling response body from create user request")
	}

	streetAddress := requestBody["street_address"]

	//CHECK IF RESTAURANT EXISTS AND PUSH IN A NEW RESTAURANT INSTANCE IF IT DOESN'T AND RETURN BLANK ARRAY
	// Seeing if the restaurant exists before trying to grab the reviews linked to it
	restaurantCheck := fmt.Sprintf("SELECT * FROM allergy_db.Restaurants WHERE street_address='%s'", streetAddress)

	res, err := db.Query(restaurantCheck)

	var restaurant RestaurantInfo

	for res.Next() {
		// for each row, scan the result into our tag composite object
		err = res.Scan(&restaurant.StreetAddress, &restaurant.Name, &restaurant.City)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// log.Printf("Queried restaurant name: %s", restaurant.Name)
	}
	if restaurant.Name == "" {
		// Case where no restaurant currently exists with the address provided
	}

	sqlRequest := fmt.Sprintf(
		`SELECT r.*
			FROM allergy_db.Restaurants R
			INNER JOIN restaurant_reviews rR
			ON rR.restaurant_address = '%s'
			INNER JOIN allergy_db.Reviews r
			ON r.id = rR.review_id`,
		streetAddress)

	// fmt.Printf("sql GET request: %s\n", sqlRequest)
	res, err = db.Query(sqlRequest)

	// fmt.Printf("Response from db Query:  | %s ", res)
	// Formatting MySQL response to JSON
	var reviews []ReviewInfo
	for res.Next() {
		// for each row, scan the result into our tag composite object
		var tempId int
		var tempReview string

		err = res.Scan(&tempId, &tempReview) // &tempId, &tempReview

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var reviewInst ReviewInfo
		reviewInst.Id = tempId
		reviewInst.Review = tempReview
		reviews = append(reviews, reviewInst)
		// log.Printf(review.Review)
	}
	// fmt.Printf("Response from db execution: %s\n", res)

	if err != nil {
		log.Println("Error pulling reviews", err)
		return responses.ServerError(err), fmt.Errorf("error pulling review from inner join table: %s", err)
	}
	defer db.Close()

	responseBody, err := json.Marshal(reviews)
	if err != nil {
		log.Println("ERROR MARSHALLING RESPONSE BODY TO JSON", err)
		return responses.ServerError(err), fmt.Errorf("ERROR MARSHALLING RESPONSE BODY TO JSON")
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 202,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseBody),
	}

	return response, nil
}
