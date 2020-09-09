package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidmm07/ordenes_api/schema"
	"github.com/graphql-go/graphql"
)

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	var s schema.Ordenes
	x := s.InitSchema()
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		type GraphQLPostBody struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		var graphQLPostBody GraphQLPostBody
		err = json.Unmarshal(body, &graphQLPostBody)
		if err != nil {
			panic(err)
		}

		token := r.Header.Get("token")

		result := graphql.Do(graphql.Params{
			Schema:         x,
			RequestString:  graphQLPostBody.Query,
			VariableValues: graphQLPostBody.Variables,
			OperationName:  graphQLPostBody.OperationName,
			Context:        context.WithValue(context.Background(), "token", token),
		})
		json.NewEncoder(w).Encode(result)
	default:
		fmt.Fprintf(w, "Lo sentimos, solo metodos POST son soportados")
	}
}
