package responses

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type MessageBody struct {
	Message interface{} `json:"message"` //default value of the message
}

const (
	BadRequestMessage  = "Bad request message"
	ServerErrorMessage = "Server error"
)

func makeResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: body,
	}
}

func makeMessageBody(bodyMessage interface{}) string {
	body := MessageBody{
		Message: bodyMessage,
	}
	bodyBytes, _ := json.Marshal(body)
	return string(bodyBytes)
}

func Ok(body []byte) events.APIGatewayProxyResponse {
	return makeResponse(http.StatusOK, string(body))
}

func Created(body []byte) events.APIGatewayProxyResponse {
	return makeResponse(http.StatusCreated, string(body))
}

func Accepted(body []byte) events.APIGatewayProxyResponse {
	return makeResponse(http.StatusAccepted, string(body))
}
