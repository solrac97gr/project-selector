build:
	@go build -o bin/project-selector main.go

install:
	@go install

uninstall:
	@rm -f $(GOPATH)/bin/project-selector

run:
	@go run main.go

clean:
	@rm -rf bin
