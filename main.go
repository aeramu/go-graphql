package main

import(
  "context"
  "encoding/json"

  "github.com/aeramu/go-graphql/resolver"

  "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
  "github.com/graph-gophers/graphql-go"
)

func main(){
  lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error){
  schema := graphql.MustParseSchema(resolver.Schema, &resolver.Resolver{})

  var params struct{
      Query string
      OperationName string
      Variables map[string]interface{}
  }
  json.Unmarshal([]byte(request.Body), &params)

  ctx1 := context.WithValue(ctx, "token", request.Headers["token"])

  response := schema.Exec(ctx1, params.Query, params.OperationName, params.Variables)
  responseJSON,_ := json.Marshal(response)

  return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}
