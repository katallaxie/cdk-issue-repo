.PHONY: test
test:
	go test -v -race ./... -short -cover -coverprofile cover.out