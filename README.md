# go-graphql

## What is it?
It's a golang back-end server for QnA app with login, register, askQuestion, and answerQuestion implemented.<br/>
This project using graphql implementation for API, and for the repository, it has two implementation: AWS DynamoDB and Firestore GCP.
For the deployment it using AWS lambda

## And then what?
It's just a learning project. I learned a lot at this project.
1. Golang, it's first time i'm using golang.
2. AWS Lambda, It's first time using it.
3. AWS Dynamo DB, it's my first time too.
4. Firestore, it's my first time.
5. **Graphql, of course** it's OP
6. **Architecture:** yeah, architecture. At this project, i learned how important architecture is. With my bad architecture, it's so hard to change implementation from AWS DynamoDB and Firestore. And then i imagined what if it's got more complicated? So, in the next project i learned about clean architecture. It's my important step to learn that clean architecture.

## Dependency

for main.go:
  - package resolver
  - "github.com/graph-gophers/graphql-go"
  - "github.com/aws/aws-lambda-go/events"
  - "github.com/aws/aws-lambda-go/lambda"

for test.go:
  - package resolver
  - "github.com/graph-gophers/graphql-go"
  - "github.com/graph-gophers/graphql-go/relay"
  - "github.com/friendsofgo/graphiql"

for resolver:
  - package service
  - package model
  - "github.com/graph-gophers/graphql-go"

for service:
  - package repository
  - package entity
  - package model
  - "github.com/graph-gophers/graphql-go"
  - "github.com/dgrijalva/jwt-go"
  - "github.com/google/uuid"

for repository:
  - "github.com/aws/aws-sdk-go/aws"
  - "github.com/aws/aws-sdk-go/aws/session"
  - "github.com/aws/aws-sdk-go/service/dynamodb"
  - "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

for model:
  - "github.com/graph-gophers/graphql-go"

for entity: none
