package service

import(
  "github.com/aeramu/go-graphql/repository"
  "github.com/aeramu/go-graphql/model"
  "github.com/aeramu/go-graphql/entity"

  "github.com/graph-gophers/graphql-go"
  "github.com/google/uuid"
)

type AccountService interface{
  GetAccountById(ID graphql.ID)(*model.AccountModel, error)
  LoginAccount(email string, username string, password string)(string, error)
  RegisterAccount(email string, username string, password string)(string, error)
}

func NewAccountService()(AccountService){
  return &AccountServiceImplementation{
    repo: repository.NewAccountRepository(),
  }
}

type AccountServiceImplementation struct{
    repo repository.AccountRepository
}

func (service *AccountServiceImplementation) GetAccountById(ID graphql.ID)(*model.AccountModel, error){
  account,_ := service.repo.GetItemById(string(ID))
  accountModel := &model.AccountModel{
    ID: graphql.ID(account.ID),
    Email: account.Email,
    Username: account.Username,
  }
  return accountModel, nil
}

func (service *AccountServiceImplementation) LoginAccount(email string, username string, password string)(string, error){
  account := new(entity.AccountEntity)

  //checking if email or username already taken
  if email != ""{
    account,_ = service.repo.GetItemByIndex("Email", email);
  } else if username != ""{
    account,_ = service.repo.GetItemByIndex("Username", username)
  } else{
    return "email and username field empty", nil
  }
  if account == nil{
      return "email or username not registered", nil
  }

  if password != account.Password{
    return "Wrong Password", nil
  }

  // create JWT to send to the client
  token := CreateJWT(account.ID)

  return token, nil
}

func (service *AccountServiceImplementation) RegisterAccount(email string, username string, password string)(string, error){
  //checking if email or username already taken
  if account,_ := service.repo.GetItemByIndex("Email", email); account != nil{
    return "email already registered"
  }
  if account,_ := service.repo.GetItemByIndex("Username", username); account != nil{
    return "username already taken"
  }

  // create account entity
  account := &entity.AccountEntity{
    ID: uuid.New().String(),
    Email: email,
    Username: username,
    Password: password,
  }

  // put account to the repository
  service.repo.PutItem(account)

  // create JWT to send to the client
  token := CreateJWT(account.ID)

  return token
}
