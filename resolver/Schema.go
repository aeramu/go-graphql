package resolver

type Resolver struct{}

var Schema = `
  schema{
    query: Query
    mutation: Mutation
  }

  type Query{
    me: Account
    account(id: ID!): Account
    question(id: ID!): Question
    questionList: QuestionConnection!
  }

  type Mutation{
    registerAccount(email: String!, username: String!, password: String!): String!
    loginAccount(email: String = "", username: String = "", password: String!): String!
    askQuestion(title: String!, body: String!): Question!
    answerQuestion(questionID: ID!, body: String!): Answer!
  }

  type Account{
    id: ID!
    email: String!
    username: String!
  }

  type Question{
    id: ID!
    title: String!
    body: String!
    answers: [Answer]!
    author: Account!
  }

  type QuestionConnection{
      edges: [QuestionEdge]!
      pageInfo: PageInfo!
  }

  type QuestionEdge{
    cursor: ID!
    node: Question
  }

  type Answer{
    id: ID!
    body: String!
    author: Account!
  }

  type PageInfo{
    startCursor: ID
    endCursor: ID
    hasNextPage: Boolean!
  }
`
