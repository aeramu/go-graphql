package resolver

import (
  "github.com/aeramu/go-graphql/entity"

  "github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

type AccountResolver struct{
  a *entity.AccountEntity
}
func (r *AccountResolver) ID()(graphql.ID){
  // return id in the form of graphql ID, not string
  return graphql.ID(r.a.ID)
}
func (r *AccountResolver) Email()(string){
  return r.a.Email
}
func (r *AccountResolver) Username()(string){
  return r.a.Username
}

type QuestionResolver struct{
  q *entity.QuestionEntity
}
func (r *QuestionResolver) ID()(graphql.ID){
  // return id in the form of graphql ID, not string
  return graphql.ID(r.q.ID)
}
func (r *QuestionResolver) Title()(string){
  return r.q.Title
}
func (r *QuestionResolver) Body()(string){
  return r.q.Body
}
func (r *QuestionResolver) Answers()([]*AnswerResolver){
  var answers []*AnswerResolver
  //  get answer resolver from every id in answers
  for _,answer := range(r.q.Answers) {
    answers = append(answers, &AnswerResolver{answer})
  }
  return answers
}
func (r *QuestionResolver) Author()(*AccountResolver){
  // get account from id
  return new(Resolver).Account(struct{ID graphql.ID}{ID: graphql.ID(r.q.Author)})
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
  // get account from id
  return new(Resolver).Account(struct{ID graphql.ID}{ID: graphql.ID(r.a.Author)})
}
