package resolver

import(
  "github.com/aeramu/go-graphql/repository"
  "github.com/aeramu/go-graphql/entity"

  "time"
  "github.com/google/uuid"
  "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) AnswerQuestion(args struct{
  QuestionID graphql.ID
  Body string
})(*AnswerResolver){
  //get accountID from jwt (TODO: JWTService and get header from request)
  accountID := "3"
  //create answer entity
  answer := &entity.AnswerEntity{
    ID: uuid.New().String(),
    Body: args.Body,
    Author: accountID,
  }
  // create question repository to access DB
  questionRepository := repository.NewQuestionRepository()
  //put answer to DB
  questionRepository.UpdateItemListAdd(string(args.QuestionID), "Answers", []*entity.AnswerEntity{answer})
  //return asnwer resolver
  return &AnswerResolver{answer}
}

func (r *Resolver) Question(args struct{
  ID graphql.ID
})(*QuestionResolver){
  // create question repository to access DB
  questionRepository := repository.NewQuestionRepository()
  // get question from DB
  question,_ := questionRepository.GetItemById(string(args.ID))
  // return questionResolver
  return &QuestionResolver{question}
}

func (r *Resolver) AskQuestion(args struct{
  Title string
  Body string
})(*QuestionResolver){
  //get accountID from jwt (TODO: JWTService and get header from request)
  accountID := "1"
  // create question entity
  question := &entity.QuestionEntity{
    ID: uuid.New().String(),
    Title: args.Title,
    Body: args.Body,
    Author: accountID,
    Type: "Question",
    Timestamp: time.Now().UnixNano(),
  }
  // create question repository to access DB
  questionRepository := repository.NewQuestionRepository()
  // put question to  DB
  questionRepository.PutItem(question)
  // return questionResolver
  return &QuestionResolver{question}
}

func (r *Resolver) QuestionList()(*QuestionConnectionResolver){
  // create question repository to access DB
  questionRepository := repository.NewQuestionRepository()
  // get question list from repository
  questionList,_ := questionRepository.GetItemListSorted("Timestamp")
  // return question conncetion resolver
  return &QuestionConnectionResolver{questionList}
}

type QuestionResolver struct{
  a *entity.QuestionEntity
}
func (r *QuestionResolver) ID()(graphql.ID){
  return graphql.ID(r.a.ID)
}
func (r *QuestionResolver) Title()(string){
  return r.a.Title
}
func (r *QuestionResolver) Body()(string){
  return r.a.Body
}
func (r *QuestionResolver) Answers()([]*AnswerResolver){
  var answers []*AnswerResolver
  //  get answer resolver from every answer in array
  for _,answer := range(r.a.Answers) {
    answers = append(answers, &AnswerResolver{answer})
  }
  return answers
}
func (r *QuestionResolver) Author()(*AccountResolver){
  input := struct{ID graphql.ID}{
    ID: graphql.ID(r.a.Author),
  }
  return new(Resolver).Account(input)
}

type AnswerResolver struct{
  a *entity.AnswerEntity
}
func (r *AnswerResolver) ID()(graphql.ID){
  return graphql.ID(r.a.ID)
}
func (r *AnswerResolver) Body()(string){
  return r.a.Body
}
func (r *AnswerResolver) Author()(*AccountResolver){
  input := struct{ID graphql.ID}{
    ID: graphql.ID(r.a.Author),
  }
  return new(Resolver).Account(input)
}

type QuestionEdgeResolver struct{
  a *entity.QuestionEntity
}
func (r *QuestionEdgeResolver) Cursor()(graphql.ID){
  return graphql.ID(r.a.ID)
}
func (r *QuestionEdgeResolver) Node()(*QuestionResolver){
  return &QuestionResolver{r.a}
}

type QuestionConnectionResolver struct{
  questionList []*entity.QuestionEntity
  lastKey string
}
func (r *QuestionConnectionResolver) Edges()([]*QuestionEdgeResolver){
  var edges []*QuestionEdgeResolver
  //  get answer edge resolver from every edge in edges
  for _,edge := range(r.questionList) {
    edges = append(edges, &QuestionEdgeResolver{edge})
  }
  return edges
}
func (r *QuestionConnectionResolver) PageInfo()(*PageInfoResolver){
  return &PageInfoResolver{r.m.PageInfo}
}


// type PageInfoResolver struct{
//   m *model.PageInfoModel
// }
// func (r *PageInfoResolver) StartCursor()(graphql.ID){
//   return r.m.StartCursor
// }
// func (r *PageInfoResolver) EndCursor()(graphql.ID){
//   return r.m.EndCursor
// }
// func (r *PageInfoResolver) HasNextPage()(bool){
//   return r.m.HasNextPage
// }
