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
    questionList(first: Int = 30, after: ID): QuestionConnection!
  }

  type Mutation{
    registerAccount(email: String!, username: String!, password: String!): String!
    loginAccount(email: String = "", username: String = "", password: String!): String!
    answerQuestion(questionID: ID!, body: String!): Answer!
    askQuestion(title: String!, body: String!): Question!
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
    node: Question!
  }

  type PageInfo{
    startCursor: ID!
    endCursor: ID!
    hasNextPage: Boolean!
  }

  type Answer{
    id: ID!
    body: String!
    author: Account!
  }
`
