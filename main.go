// main.go
package main

import (
	"user-crud/config"
	"user-crud/controllers"
	"user-crud/utils"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	config.ConnectDB()
}

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	httpReq, err := utils.ConvertToHTTPRequest(request)
	if err != nil {
		return events.LambdaFunctionURLResponse{StatusCode: 500}, nil
	}

	rec := utils.NewRecorder()

	path := request.RequestContext.HTTP.Path
	method := request.RequestContext.HTTP.Method
	log.Println("path ", path, " method ", method)
	switch path {

	case "/users":
		switch method {
		case "GET":
			controllers.GetUsers(rec, httpReq)
		case "POST":
			controllers.CreateUser(rec, httpReq)
		case "PUT":
			controllers.UpdateUser(rec, httpReq)
		case "DELETE":
			controllers.DeleteUser(rec, httpReq)
		}
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: rec.StatusCode,
		Body:       rec.Body.String(),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
