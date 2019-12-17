package entity

type AccountEntity struct{
  ID string
  Email string
  Username string
  Password string
}

type QuestionEntity struct{
  ID string
  Title string
  Body string
  Answers []*AnswerEntity // array of answer id
  Author string // account id
  Type string
  Timestamp int64
}

type AnswerEntity struct{
  ID string
  Body string
  Author string // account id
}
