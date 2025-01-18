package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nalis-code/consumer-service/lambdas/pricedata/internal"
)

func main() {
	h := internal.NewHandler()
	lambda.Start(h.Handle)
}
