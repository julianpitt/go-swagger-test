generate-swagger:
	swagger mixin --format=yaml --output swagger.yaml ./spec/base.yaml ./spec/**/*.yaml

generate-server:
	swagger generate server --name backend --model-package restapi/models
	go build -o server ./cmd/backend-server/main.go

generate-client:
	swagger generate client --name=HelloWorld --additional-initialism=cls --target=sdk/go
	go build -o client ./cmd/client-test/main.go

generate: generate-swagger generate-server generate-client