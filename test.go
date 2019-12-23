package main

import(
  "github.com/aeramu/go-graphql/resolver"

  "net/http"
  "fmt"

  graphql "github.com/graph-gophers/graphql-go"
  "github.com/graph-gophers/graphql-go/relay"
  "github.com/friendsofgo/graphiql"
)

func main(){
  fmt.Println("running on localhost:3000")

  schema := graphql.MustParseSchema(resolver.SchemaTest, &resolver.Resolver{})
  http.Handle("/query",&relay.Handler{Schema: schema})

  // graphiql
  graqhiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
  if err != nil{
    panic(err)
  }

  http.Handle("/",graqhiqlHandler)
  http.ListenAndServe(":3000", nil)

}
