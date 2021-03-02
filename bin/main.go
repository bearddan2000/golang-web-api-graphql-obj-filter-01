package main

import (

  "encoding/json"

  "fmt"

  "log"

  "strings"

  "github.com/graphql-go/graphql"
)

type Country struct {
  Abbr string
  Name string
}

var countryType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Country",
        Fields: graphql.Fields{
            "Abbr": &graphql.Field{
                Type: graphql.String,
            },
            "Name": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

func populate() []Country {

    countries := []Country {
      {Abbr: "ag", Name: "Argentina"},
      {Abbr: "au", Name: "Australia"},
      {Abbr: "be", Name: "Belgium"},
      {Abbr: "br", Name: "Brazil"},
      {Abbr: "ca", Name: "Canada"},
      {Abbr: "mx", Name: "Mexico"},
      {Abbr: "cu", Name: "Cuba"},
      {Abbr: "nl", Name: "Netherlands"},
      {Abbr: "en", Name: "Britian"},
      {Abbr: "de", Name: "Germany"},
    }

    return countries
}
func main() {

    countries := populate()

    // Schema
    fields := graphql.Fields{
        "Countries": &graphql.Field{
            Type: graphql.NewList(countryType),
            Args: graphql.FieldConfigArgument{
                "startWith": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
              results := make([]Country, 0)
              id, ok := p.Args["startWith"].(string)
                if ok {
                    // Parse our tutorial array for the matching id
                    for _, item := range countries {
                        if strings.HasPrefix(item.Abbr, id) {
                            // return our tutorial
                            results = append(results, item)
                        }
                    }
                }
                return results, nil
            },
        },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query := `
        {
          Countries(startWith: "b") {
            Abbr
            Name
          }
        }
    `
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
