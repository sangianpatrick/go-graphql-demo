package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postQuery struct {
	Query string `json:"query"`
}

// GraphqlHandler is a concrete struct of graphql handler.
type GraphqlHandler struct {
	Schema graphql.Schema
}

// NewGraphqlHandler is a constructor.
func NewGraphqlHandler(mux *http.ServeMux, schema graphql.Schema) {
	handler := &GraphqlHandler{
		Schema: schema,
	}

	mux.HandleFunc("/graphql-demo", handler.Handle)
}

// Handle will hanlde all graphql incoming request.
func (h GraphqlHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var result *graphql.Result
	var requestString string
	var ctx context.Context

	ctx = h.wrapContext(r)
	requestString = h.getRequestString(r)

	result = graphql.Do(graphql.Params{
		Schema:        h.Schema,
		RequestString: requestString,
		Context:       ctx,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h GraphqlHandler) wrapContext(r *http.Request) (ctx context.Context) {
	return r.Context()
}

func (h GraphqlHandler) getRequestString(r *http.Request) (requestString string) {
	requestString = r.URL.Query().Get("query")

	if requestString == "" {
		var pq postQuery
		json.NewDecoder(r.Body).Decode(&pq)
		requestString = pq.Query
	}

	return
}
