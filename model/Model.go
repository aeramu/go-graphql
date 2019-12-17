package model

import(
  "github.com/graph-gophers/graphql-go"
)

type AccountModel struct{
  ID graphql.ID
  Email string
  Username string
}

type QuestionModel struct{
  ID graphql.ID
  Title string
  Body string
  Answers []*AnswerModel
  Author *AccountModel
}

type QuestionEdgeModel struct{
  Cursor graphql.ID
  Node *QuestionModel
}

type QuestionConnectionModel struct{
  Edges []*QuestionEdgeModel
  PageInfo *PageInfoModel
}

type AnswerModel struct{
  ID graphql.ID
  Body string
  Author *AccountModel
}

type PageInfoModel struct{
  StartCursor graphql.ID
  EndCursor graphql.ID
  HasNextPage bool
}
