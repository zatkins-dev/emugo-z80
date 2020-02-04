GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME_WIN=hw1-atkins.exe
BINARY_NAME=hw1-atkins

all:
	$(GOBUILD) -o $(BINARY_NAME_WIN)
	$(GOBUILD) -o $(BINARY_NAME)

clean:
	$(GOCLEAN)

run:
	$(GORUN) main.go

dist: tar

tar:
	git archive --format=tar -o ../$(BINARY_NAME).tar HEAD^{tree}

zip:
	git archive --format=zip -o ../$(BINARY_NAME).zip HEAD^{tree}