# GaneshaHub

##Dependency

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
