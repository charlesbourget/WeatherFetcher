NAME=WeatherFetcher
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Run in development mode.
run: build
	@go run . -- -e development

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@go build -o $(NAME)
	@GIN_MODE=release ./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: format
## test: Run gofmt on all source files
format:
	@gofmt -s -w .

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
