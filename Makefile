GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_WIN=hw1-atkins.exe
BINARY_NAME=z80-atkins

all:
	$(GOBUILD) -o $(BINARY_NAME)
	$(GOBUILD) -o $(BINARY_NAME_WIN)

clean:
	$(GOCLEAN)
	rm $(BINARY_NAME)
	rm $(BINARY_NAME_WIN)

run:
	$(GORUN) main.go
