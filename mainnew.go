package main

import (
	"context"

	"user-crud/config"
	"user-crud/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var adapter *httpadapter.HandlerAdapter

func init() {
	config.ConnectDB()
	adapter = httpadapter.New(routes.RegisterRoutes())
}

func handlernew(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func mainnew() {
	lambda.Start(handlernew)
}
