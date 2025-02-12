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

config:
	@mkdir -p $(HOME)/.config/project-selector
	@cp config.json $(HOME)/.config/project-selector/config.json
	@nano $(HOME)/.config/project-selector/config.json


.PHONY: build install uninstall run clean config open-config
