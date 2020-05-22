package firestoreRepository

import (
	"github.com/aeramu/go-graphql/entity"
  
	"context"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// interface for account repository
type AccountRepository interface{
	PutItem(account *entity.AccountEntity) (error)
	//GetItemById(ID string) (*entity.AccountEntity, error)
	//GetItemByIndex(indexName string, indexValue string) (*entity.AccountEntity, error)
}

// Constructor for AccountRepository
func NewAccountRepository()(AccountRepository){
	ctx := context.Background()
	credential := option.WithCredentialsFile("credential.json")
	app,_ := firebase.NewApp(ctx, nil, credential)
	client,_ := app.Firestore(ctx)
	return &AccountRepositoryImplementation{
		client: client,
	}
}

// Class for account repository implementation
type AccountRepositoryImplementation struct{
	client *firestore.Client
}

func (repository *AccountRepositoryImplementation) PutItem(account *entity.AccountEntity) (error){
	ctx := context.Background()
	_, _, err := repository.client.Collection("account").Add(ctx, map[string]interface{}{
        "Email": account.Email,
        "Username":  account.Username,
        "Password":  account.Password,
	})
	if err != nil {
		return err
	}
	return nil
}