swagger:
	swagger generate spec -o ./swagger.yml -m -w ./api && swagger generate server -f ./swagger.yml --exclude-main -t . -A case-mgmt --existing-models=github.com/case-management-suite/models -a ops --skip-tag-packages  --server-package=restapi

mocks:
	mockgen -package mocks -destination=mocks/mock_runtime_producer.go github.com/go-openapi/runtime Producer 
	mockgen -source=controllers/case-controller-api.go -destination=mocks/mock_case_controller_api.go -package=mocks

gql:
	cd graph
	go run github.com/99designs/gqlgen generate