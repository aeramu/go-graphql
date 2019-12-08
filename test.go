package main

import(
  "github.com/aeramu/go-graphql/resolver"

  "net/http"

  graphql "github.com/graph-gophers/graphql-go"
  "github.com/graph-gophers/graphql-go/relay"
  "github.com/friendsofgo/graphiql"
)

func main(){
  schema := graphql.MustParseSchema(resolver.Schema, &resolver.Resolver{})
  http.Handle("/query",&relay.Handler{Schema: schema})

  // graphiql
  graqhiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
  if err != nil{
    panic(err)
  }

  http.Handle("/",graqhiqlHandler)
  http.ListenAndServe(":3000", nil)

}
