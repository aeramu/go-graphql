package resolver

var SchemaTest = `
  schema{
    query: Query
    mutation: Mutation
  }

  type Query{
    account(id: ID!): Account
  }

  type Mutation{
    registerAccount(email: String!, username: String!, password: String!): String!
    loginAccount(email: String = "", username: String = "", password: String!): String!
    answerQuestion(questionID: ID!, body: String!): Answer!
  }

  type Account{
    id: ID!
    email: String!
    username: String!
  }

  type Answer{
    id: ID!
    body: String!
    author: Account!
  }
`
