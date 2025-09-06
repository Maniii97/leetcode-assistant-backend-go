package main

import (
	"context"
	"lc-assistant-backend/internal/config"
	"lc-assistant-backend/internal/routes"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Println("Initializing Lambda function...")
	config.Load()
	router := routes.SetupRoutes()
	ginLambda = ginadapter.New(router)
	log.Println("Lambda function initialized successfully")
}

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	log.Printf("Function URL request: %s %s", req.RequestContext.HTTP.Method, req.RawPath)

	// Convert Function URL request to API Gateway format
	apiGWReq := events.APIGatewayProxyRequest{
		HTTPMethod:            req.RequestContext.HTTP.Method,
		Path:                  req.RawPath,
		QueryStringParameters: req.QueryStringParameters,
		Headers:               req.Headers,
		Body:                  req.Body,
		IsBase64Encoded:       req.IsBase64Encoded,
		RequestContext: events.APIGatewayProxyRequestContext{
			HTTPMethod: req.RequestContext.HTTP.Method,
			Path:       req.RawPath,
		},
	}

	// Process with gin
	apiGWResp, err := ginLambda.ProxyWithContext(ctx, apiGWReq)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Internal server error"}`,
		}, nil
	}

	// Convert back to Function URL response
	return events.LambdaFunctionURLResponse{
		StatusCode:      apiGWResp.StatusCode,
		Headers:         apiGWResp.Headers,
		Body:            apiGWResp.Body,
		IsBase64Encoded: apiGWResp.IsBase64Encoded,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
