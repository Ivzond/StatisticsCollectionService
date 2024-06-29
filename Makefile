GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

BINARY_NAME = statistics-collection-service

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/StatisticsCollectionService/main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/StatisticsCollectionService/main.go
	./$(BINARY_NAME)

.PHONY: all build test clean run
