package resolver

import (
  "github.com/aeramu/go-graphql/model"
  "github.com/graph-gophers/graphql-go"
)

type AccountResolver struct{
  m *model.AccountModel
}
func (r *AccountResolver) ID()(graphql.ID){
  return graphql.ID(r.m.ID)
}
func (r *AccountResolver) Email()(string){
  return r.m.Email
}
func (r *AccountResolver) Username()(string){
  return r.m.Username
}

type QuestionResolver struct{
  m *model.QuestionModel
}
func (r *QuestionResolver) ID()(graphql.ID){
  return graphql.ID(r.m.ID)
}
func (r *QuestionResolver) Title()(string){
  return r.m.Title
}
func (r *QuestionResolver) Body()(string){
  return r.m.Body
}
func (r *QuestionResolver) Answers()(*AnswerConnectionResolver){
  return &AnswerConnectionResolver{r.m.Answers}
}
func (r *QuestionResolver) Author()(*AccountResolver){
  return &AccountResolver{r.m.Author}
}

type QuestionEdgeResolver struct{
  m *model.QuestionEdgeModel
}
func (r *QuestionEdgeResolver) Cursor()(graphql.ID){
  return r.m.Cursor
}
func (r *QuestionEdgeResolver) Node()(*QuestionResolver){
  return &QuestionResolver{r.m.Node}
}

type QuestionConnectionResolver struct{
  m *model.QuestionConnectionModel
}
func (r *QuestionConnectionResolver) Edges()([]*QuestionEdgeResolver){
  var edges []*QuestionEdgeResolver
  //  get answer edge resolver from every edge in edges
  for _,edge := range(r.m.Edges) {
    edges = append(edges, &QuestionEdgeResolver{edge})
  }
  return edges
}
func (r *QuestionConnectionResolver) PageInfo()(*PageInfoResolver){
  return &PageInfoResolver{r.m.PageInfo}
}

type AnswerResolver struct{
  m *model.AnswerModel
}
func (r *AnswerResolver) ID()(graphql.ID){
  return graphql.ID(r.m.ID)
}
func (r *AnswerResolver) Body()(string){
  return r.m.Body
}
func (r *AnswerResolver) Author()(*AccountResolver){
  return &AccountResolver{r.m.Author}
}

type AnswerEdgeResolver struct{
  m *model.AnswerEdgeModel
}
func (r *AnswerEdgeResolver) Cursor()(graphql.ID){
  return r.m.Cursor
}
func (r *AnswerEdgeResolver) Node()(*AnswerResolver){
  return &AnswerResolver{r.m.Node}
}

type AnswerConnectionResolver struct{
  m *model.AnswerConnectionModel
}
func (r *AnswerConnectionResolver) Edges()([]*AnswerEdgeResolver){
  var edges []*AnswerEdgeResolver
  //  get answer edge resolver from every edge in edges
  for _,edge := range(r.m.Edges) {
    edges = append(edges, &AnswerEdgeResolver{edge})
  }
  return edges
}
func (r *AnswerConnectionResolver) PageInfo()(*PageInfoResolver){
  return &PageInfoResolver{r.m.PageInfo}
}

type PageInfoResolver struct{
  m *model.PageInfoModel
}
func (r *PageInfoResolver) StartCursor()(graphql.ID){
  return r.m.StartCursor
}
func (r *PageInfoResolver) EndCursor()(graphql.ID){
  return r.m.EndCursor
}
func (r *PageInfoResolver) HasNextPage()(bool){
  return r.m.HasNextPage
}
