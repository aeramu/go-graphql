package resolver

import(
  "github.com/aeramu/go-graphql/service"

  "github.com/graph-gophers/graphql-go"
)

var SchemaTest = `
  schema{
    query: Query
    mutation: Mutation
  }

  type Query{
    account(id: ID!): Account
  }

  type Mutation{
    registerAccount(email: String!, username: String!, password: String!): String!
    loginAccount(email: String = "", username: String = "", password: String!): String!
  }

  type Account{
    id: ID!
    email: String!
    username: String!
  }
`

type Resolver struct{}

//Query
func (r *Resolver) Account(args struct{ID graphql.ID})(*AccountResolver){
  service := service.NewAccountService()
  account,_ := service.GetAccountById(args.ID)
  return &AccountResolver{account}
}

func (r *Resolver) LoginAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  service := service.NewAccountService()
  token,_ := service.LoginAccount(args.Email,args.Username,args.Password)
  return token
}

func (r *Resolver) RegisterAccount(args struct{
  Email string
  Username string
  Password string
})(string){
  service := service.NewAccountService()
  token,_ := service.RegisterAccount(args.Email,args.Username,args.Password)
  return token
}
// func (r *Resolver) Me()(*AccountResolver){
//   service := service.NewAccountService()
//   account := service.GetAccountById()
//   return &AccountResolver{account}
// }
//
// func (r *Resolver) Question(args struct{ID graphql.ID})(*QuestionResolver){
//   service := service.NewQnAService()
//   question := service.GetQuestionById(args.ID)
//   return &QuestionResolver{question}
// }
//
// func (r *Resolver) QuestionList()(*QuestionConnectionResolver){
//   service := service.NewQnAService()
//   questionConnection := service.GetQuestionConnection()
//   return &QuestionConnectionResolver{questionConnection}
// }
//
// func (r *Resolver) AskQuestion(args struct{
//   Title string
//   Body string
// })(*QuestionResolver){
//   return &QuestionResolver{question}
// }
//
// func (r *Resolver) AnswerQuestion(args struct{
//   QuestionID graphql.ID
//   Body string
// })(*AnswerResolver){
//   return &AnswerResolver{answer}
// }
