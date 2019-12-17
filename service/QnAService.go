package service

import(
  "github.com/aeramu/go-graphql/repository"
  "github.com/aeramu/go-graphql/model"
  "github.com/aeramu/go-graphql/entity"

  "github.com/graph-gophers/graphql-go"
  "github.com/google/uuid"
)

type QnAService interface{
  AnswerQuestion(questionID graphql.ID, body string)(*model.AnswerModel, error)
}

func NewQnAService()(QnAService){
  return &QnAServiceImplementation{
    questionRepo: repository.NewQuestionRepository(),
    accountRepo: repository.NewAccountRepository(),
  }
}

type QnAServiceImplementation struct{
  questionRepo repository.QuestionRepository
  accountRepo repository.AccountRepository
}

func (service *QnAServiceImplementation) AnswerQuestion(questionID graphql.ID, body string)(*model.AnswerModel, error){
  //get accountID from jwt (TODO: JWTService and get header from request)
  //create answer entity
  accountID := "3"
  answer := &entity.AnswerEntity{
    ID: uuid.New().String(),
    Body: body,
    Author: accountID,
  }
  //put answer in repo
  service.questionRepo.UpdateItemListAdd(string(questionID), "Answers", []*entity.AnswerEntity{answer})
  //get account of answer author for model
  account,_ := service.accountRepo.GetItemById(accountID)
  //create answer model
  answerModel := &model.AnswerModel{
    ID: graphql.ID(answer.ID),
    Body: answer.Body,
    Author: &model.AccountModel{
      ID: graphql.ID(account.ID),
      Email: account.Email,
      Username: account.Username,
    },
  }
  return answerModel, nil
}
