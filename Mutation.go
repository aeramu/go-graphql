package resolver

import(
  "github.com/aeramu/go-graphql/entity"
  "github.com/aeramu/go-graphql/repository"
  "github.com/aeramu/go-graphql/jwt"

  "github.com/graph-gophers/graphql-go"
  "github.com/google/uuid"

  "time"
)

func (r *Resolver) RegisterAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  // create account repository
  repo := repository.NewAccountRepository()

  //checking if email or username already taken
  if account,_ := repo.GetItemByIndex("Email", args.Email); account != nil{
    return "email already registered"
  }
  if account,_ := repo.GetItemByIndex("Username", args.Username); account != nil{
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
  repo.PutItem(account)

  // create JWT to send to the client
  token := jwt.CreateJWT(account.ID)

  return token
}

func (r *Resolver) LoginAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  account := new(entity.AccountEntity)
  // create account repository
  repo := repository.NewAccountRepository()

  //checking if email or username already taken
  if args.Email != ""{
    account,_ = repo.GetItemByIndex("Email", args.Email);
  } else if args.Username != ""{
    account,_ = repo.GetItemByIndex("Username", args.Username)
  } else{
    return "email and username field empty"
  }
  if account == nil{
      return "email or username not registered"
  }
  if args.Password != account.Password{
    return "Wrong Password"
  }
  // create JWT to send to the client
  token := jwt.CreateJWT(account.ID)

  return token
}

func (r *Resolver) AskQuestion(args struct{
  Title string
  Body string
})(*QuestionResolver){
  // create question entity
  accountID := "1"
  question := &entity.QuestionEntity{
    ID: uuid.New().String(),
    Title: args.Title,
    Body: args.Body,
    Author: accountID,
    Type: "Question",
    Timestamp: time.Now().UnixNano(),
  }

  // put question to repository
  repo := repository.NewQuestionRepository()
  repo.PutItem(question)

  return &QuestionResolver{question}
}

func (r *Resolver) AnswerQuestion(args struct{
  QuestionID graphql.ID
  Body string
})(*AnswerResolver){
  // create answer entity
  accountID := "3"
  answer := &entity.AnswerEntity{
    ID: uuid.New().String(),
    Body: args.Body,
    Author: accountID,
  }

  // update question's answer in repo
  repo := repository.NewQuestionRepository()
  repo.UpdateItemListAdd(string(args.QuestionID), "Answers", []*entity.AnswerEntity{answer})

  return &AnswerResolver{answer}
}
