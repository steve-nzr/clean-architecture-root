# Go targets
.PHONY: go.tests go.build go.lint
go.tests:
	@go test ./...

go.build:
	go build -o app main.go

go.lint:
	@golangci-lint run ./...

# Kubernetes targets
.PHONY: kube.dev
kube.dev:
	@skaffold dev
