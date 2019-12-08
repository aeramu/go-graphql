package resolver

import(
  //"github.com/aeramu/go-graphql/entity"
  "github.com/aeramu/go-graphql/repository"

  "github.com/graph-gophers/graphql-go"

  "fmt"
)

func (r *Resolver) Account(args struct{ID graphql.ID})(*AccountResolver){
  // get account by id from repository
  repo := repository.NewAccountRepository()
  account,_ := repo.GetItemById(string(args.ID))

  return &AccountResolver{account}
}

func (r *Resolver) Question(args struct{ID graphql.ID})(*QuestionResolver){
  // get question from repo
  repo := repository.NewQuestionRepository()
  question,_ := repo.GetItemById(string(args.ID))

  return &QuestionResolver{question}
}

func (r *Resolver) QuestionList()([]*QuestionResolver){
  // get question list from repository
  repo := repository.NewQuestionRepository()
  questionList,err := repo.GetItemListSorted("Timestamp")
  if err!= nil{
    fmt.Println(err.Error())
  }

  // create question resolver for each question
  var questionResolverList []*QuestionResolver
  for _,question := range(questionList){
    questionResolverList = append(questionResolverList, &QuestionResolver{question})
  }

  return questionResolverList
}
