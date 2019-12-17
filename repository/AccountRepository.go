package repository

import (
  "github.com/aeramu/go-graphql/entity"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
// interface for account repository
type AccountRepository interface{
  PutItem(account *entity.AccountEntity) (error)
  GetItemById(ID string) (*entity.AccountEntity, error)
  GetItemByIndex(indexName string, indexValue string) (*entity.AccountEntity, error)
}


// Constructor for AccountRepository
func NewAccountRepository()(AccountRepository){
  return &AccountRepositoryImplementation{
    // name of the table in aws dynamodb
    tableName: aws.String("AccountTable"),
    // configuration for dynamodb
    db: dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1")),
  }
}

// Class for account repository implementation
type AccountRepositoryImplementation struct{
  tableName *string
  db *dynamodb.DynamoDB
}

func (repository *AccountRepositoryImplementation) PutItem(account *entity.AccountEntity) (error){
  // convert the account object to aws dynamodb format
  item, err := dynamodbattribute.MarshalMap(account)
  if err != nil {
    return err
  }

  // create the input for query
  input := &dynamodb.PutItemInput{
    TableName: repository.tableName,
    Item: item,
  }

  // query to dynamodb
  _, err = repository.db.PutItem(input)
  if err != nil{
    return err
  }

  // return  nil if there is no error
  return nil
}


func (repository *AccountRepositoryImplementation) GetItemById(ID string) (*entity.AccountEntity, error){
  // create the input for query
  input:=&dynamodb.GetItemInput{
    TableName: repository.tableName,
    Key: map[string]*dynamodb.AttributeValue{
      "ID":{
        S: aws.String(ID),
      },
    },
  }

  // query to dynamodb
  result, err := repository.db.GetItem(input)
  if err != nil{
    return nil, err
  }

  // if there is no item found, return nil
  if result.Item == nil{
    return nil, nil
  }

  // if item found, convert from dynamodb format to account object
  account := new(entity.AccountEntity)
  err = dynamodbattribute.UnmarshalMap(result.Item, account)
  if err != nil{
    return nil, err
  }

  // return account object
  return account, nil
}

func (repository *AccountRepositoryImplementation) GetItemByIndex(indexName string, indexValue string) (*entity.AccountEntity, error){
  // create the query input
  input := &dynamodb.QueryInput{
    TableName: repository.tableName,
    IndexName: aws.String(indexName),
    ExpressionAttributeNames: map[string]*string{
      "#indexName": aws.String(indexName),
    },
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":indexValue":{
        S: aws.String(indexValue),
      },
    },
    KeyConditionExpression: aws.String("#indexName = :indexValue"),
  }

  // query the input to dynamodb
  result, err := repository.db.Query(input)
  if err != nil{
    return nil, err
  }

  // if there is item found, return nil
  if *result.Count == 0{
    return nil, nil
  }

  // if item found, convert from dynamodb format to account object
  account := new(entity.AccountEntity)
  err = dynamodbattribute.UnmarshalMap(result.Items[0], account)
  if err != nil{
    return nil, err
  }

  // return account object
  return account, nil
}
