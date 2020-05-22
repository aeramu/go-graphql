package firestoreRepository

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