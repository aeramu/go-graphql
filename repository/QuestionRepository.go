package repository

import (
  "github.com/aeramu/go-graphql/entity"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
// interface for account repository
type QuestionRepository interface{
  PutItem(question *entity.QuestionEntity) (error)
  GetItemById(ID string) (*entity.QuestionEntity, error)
  GetItemListSorted(indexName string, limit int64, startKey *entity.QuestionCursor) ([]*entity.QuestionEntity, *entity.QuestionCursor, error)
  UpdateItemListAdd(ID string, attributeName string, attributeValue interface{}) (error)
}


// Constructor for AccountRepository
func NewQuestionRepository()(QuestionRepository){
  return &QuestionRepositoryImplementation{
    // name of the table in aws dynamodb
    tableName: aws.String("QnATable"),
    // configuration for dynamodb
    db: dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1")),
  }
}

// Class for account repository implementation
type QuestionRepositoryImplementation struct{
  tableName *string
  db *dynamodb.DynamoDB
}

func (repository *QuestionRepositoryImplementation) PutItem(question *entity.QuestionEntity) (error){
  // convert the account object to aws dynamodb format
  item, err := dynamodbattribute.MarshalMap(question)
  if err != nil {
    return err
  }

  // set answers to empty list, because that fucking hell bullshit aws convert go's empty list to null
  emptyList := make([]*dynamodb.AttributeValue, 0)
  item["Answers"] = &dynamodb.AttributeValue{L: emptyList}

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


func (repository *QuestionRepositoryImplementation) GetItemById(ID string) (*entity.QuestionEntity, error){
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
  question := new(entity.QuestionEntity)
  err = dynamodbattribute.UnmarshalMap(result.Item, question)
  if err != nil{
    return nil, err
  }

  // return account object
  return question, nil
}

func (repository *QuestionRepositoryImplementation) GetItemListSorted(indexName string, limit int64, startKey *entity.QuestionCursor) ([]*entity.QuestionEntity, *entity.QuestionCursor, error){
  expressionAttributeNames := make(map[string]*string)
  expressionAttributeNames["#key"] = aws.String("Type")

  expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
  expressionAttributeValues[":value"],_ = dynamodbattribute.Marshal("Question")

  key := make(map[string]*dynamodb.AttributeValue)
  if startKey != nil{
    key,_ = dynamodbattribute.MarshalMap(startKey)
  } else{
    key = nil
  }

  input := &dynamodb.QueryInput{
    TableName: repository.tableName,
    IndexName: aws.String(indexName),
    ExpressionAttributeNames: expressionAttributeNames,
    ExpressionAttributeValues: expressionAttributeValues,
    KeyConditionExpression: aws.String("#key = :value"),
    ScanIndexForward: aws.Bool(false),
    Limit: aws.Int64(limit),
    ExclusiveStartKey: key,
  }

  result, err := repository.db.Query(input)
  if err != nil{
    return nil, nil, err
  }

  questions := make([]*entity.QuestionEntity,limit)
  err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &questions)
  if err != nil{
    return nil, nil, err
  }
  
  lastKey := new(entity.QuestionCursor)
  if result.LastEvaluatedKey != nil{
    dynamodbattribute.UnmarshalMap(result.LastEvaluatedKey, lastKey)
  } else{
    lastKey = nil
  }


  return questions, lastKey, nil
}

func (repository *QuestionRepositoryImplementation) UpdateItemListAdd(ID string, attributeName string, attributeValue interface{}) (error){
  var err error

  // make attribute name input
  expressionAttributeNames := make(map[string]*string)
  expressionAttributeNames["#key"] = aws.String(attributeName)

  // convert from object to dynamodb format
  expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
  expressionAttributeValues[":value"], err = dynamodbattribute.Marshal(attributeValue)
  if err != nil{
    return err
  }

  // create the query input
  input := &dynamodb.UpdateItemInput{
    TableName: repository.tableName,
    Key: map[string]*dynamodb.AttributeValue{
      "ID":{
        S: aws.String(ID),
      },
    },
    ExpressionAttributeValues: expressionAttributeValues,
    ExpressionAttributeNames : expressionAttributeNames,
    UpdateExpression: aws.String("set #key = list_append(#key,:value)"),
  }

  // query the input to dynamodb
  _, err = repository.db.UpdateItem(input)
  if err != nil{
    return err
  }

  // if there is no error, return nil
  return nil
}
