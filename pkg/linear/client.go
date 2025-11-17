package linear

import (
	"net/http"

	graphql "github.com/hasura/go-graphql-client"
)

type Linear struct {
	ApiKey string
}

func (l Linear) getClient() *graphql.Client {
	url := "https://api.linear.app/graphql"

	client := graphql.NewClient(url, http.DefaultClient).
		WithRequestModifier(func(r *http.Request) {
			r.Header.Set("Authorization", l.ApiKey)
		})

	return client
}
