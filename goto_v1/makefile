# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_FOLDER=build
BINARY_NAME=goto

all: build

build: 
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_FOLDER)/$(BINARY_NAME)