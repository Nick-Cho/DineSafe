package main

import (
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct{}

func (h *Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody map[string]string
}
