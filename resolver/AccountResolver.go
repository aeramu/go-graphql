package resolver

import(
  "github.com/aeramu/go-graphql/repository"
  "github.com/aeramu/go-graphql/entity"
  "github.com/aeramu/go-graphql/service"

  "context"
  //"encoding/json"
  "github.com/google/uuid"
  "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Account(args struct{
  ID graphql.ID
})(*AccountResolver){
  // create account repository to access DB
  accountRepository := repository.NewAccountRepository()
  // get account from DB
  account,_ := accountRepository.GetItemById(string(args.ID))
  // return accountResolver
  return &AccountResolver{account}
}

func (r *Resolver) LoginAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  // create account repository to access DB
  accountRepository := repository.NewAccountRepository()
  // create empty account entity
  account := new(entity.AccountEntity)
  // get account entity from DB with email or username
  if args.Email != ""{
    account,_ = accountRepository.GetItemByIndex("Email", args.Email);
  } else if args.Username != ""{
    account,_ = accountRepository.GetItemByIndex("Username", args.Username)
  } else{
    return "email and username field empty"
  }
  // if account not found in DB
  if account == nil{
      return "email or username not registered"
  }
  // checking password
  if args.Password != account.Password{
    return "Wrong Password"
  }
  // create JWT to send to the client
  token := service.CreateJWT(account.ID)
  // send jwt to client
  return token
}

func (r *Resolver) RegisterAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  // create account repository to access DB
  accountRepository := repository.NewAccountRepository()
  //checking if email or username already taken
  if account,_ := accountRepository.GetItemByIndex("Email", args.Email); account != nil{
    return "email already registered"
  }
  if account,_ := accountRepository.GetItemByIndex("Username", args.Username); account != nil{
    return "username already taken"
  }
  // create account entity
  account := &entity.AccountEntity{
    ID: uuid.New().String(),
    Email: args.Email,
    Username: args.Username,
    Password: args.Password,
  }
  // put account to the repository
  accountRepository.PutItem(account)
  // create JWT to send to the client
  token := service.CreateJWT(account.ID)
  // return jwt to the client
  return token
}

func (r *Resolver) Me(ctx context.Context)(*AccountResolver){
  token := ctx.Value("token").(string)
	accountID := service.DecodeJWT(token)
  // create account repository to access DB
  accountRepository := repository.NewAccountRepository()
  // get account from DB
  account,_ := accountRepository.GetItemById(accountID)
  // return accountResolver
  return &AccountResolver{account}
}

type AccountResolver struct{
  a *entity.AccountEntity
}
func (r *AccountResolver) ID()(graphql.ID){
  return graphql.ID(r.a.ID)
}
func (r *AccountResolver) Email()(string){
  return r.a.Email
}
func (r *AccountResolver) Username()(string){
  return r.a.Username
}
