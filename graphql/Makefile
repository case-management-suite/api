gen:
	cd graph
	go run github.com/99designs/gqlgen generate

testGen:
	cd test
	cp graph/schema.graphqls test/schema.graphql
	go run github.com/Khan/genqlient test/genqlient.yaml