NAME=btm
TAG=latest
IMAGE_NAME=$(NAME):$(TAG)
OUT=bin/$(NAME)

.PHONY: install-deps
install-deps:
	@go get -d ./...

.PHONY: lint
lint:
	@golint ./...

.PHONY: vet
vet:
	 @go vet ./...

.PHONY: build
build:
	 @go build -o $(OUT)  ./...

.PHONY: test
test: vet
	@go test ./...

.PHONY: install
install: test
	 @go install ./...

.PHONY: docker-image
docker-image:
	@docker build -t $(IMAGE_NAME) .
