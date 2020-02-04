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
	tar -czvf --exclude-vcs ../$(BINARY_NAME).tar.gz ../$(BINARY_NAME)