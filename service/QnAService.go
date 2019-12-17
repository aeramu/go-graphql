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
  GetQuestionById(ID graphql.ID)(*model.QuestionModel, error)
}

func NewQnAService()(QnAService){
  return &QnAServiceImplementation{
    QuestionRepository: repository.NewQuestionRepository(),
  }
}

type QnAServiceImplementation struct{
  QuestionRepository repository.QuestionRepository
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
  service.QuestionRepository.UpdateItemListAdd(string(questionID), "Answers", []*entity.AnswerEntity{answer})
  //get accountModel of answer's author
  accountService := NewAccountService()
  accountModel,_ := accountService.GetAccountById(accountID)
  //create answer model
  answerModel := &model.AnswerModel{
    ID: graphql.ID(answer.ID),
    Body: answer.Body,
    Author: accountModel,
  }
  return answerModel, nil
}

func (service *QnAServiceImplementation) GetQuestionById(ID graphql.ID)(*model.QuestionModel, error){
  //get question entity from repo
  question,_ := service.QuestionRepository.GetItemById(ID)
  //create account service to get accountModel of author
  accountService := NewAccountService()
  //convert answersEntity to answersModel
  var answersModel []*model.AnswerModel
  for _,answer := range(question.Answers) {
    //get accountModel of answer author
    accountModel := accountService.GetAccountById(answer.Author)
    answerModel := &model.AnswerModel{
      ID: answer.ID,
      Body: answer.Body,
      Author: accountModel,
    }
    answersModel = append(answersModel, answerModel)
  }
  //get accountModel of question author
  accountModel,_ := accountService.GetAccountById(question.Author)
  //parse entity to model
  questionModel := &model.QuestionModel{
    ID: question.ID,
    Title: question.Title,
    Body: question.Title,
    Answers: answersModel,
    Author: accountModel,
  }
}
