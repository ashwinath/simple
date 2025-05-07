.PHONY: test
test:
	@go clean -testcache
	@go test -v -coverprofile cover.out -race ./...
	@go tool cover -func cover.out
	@rm cover.out
