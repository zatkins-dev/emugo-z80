GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME_WIN=hw1-atkins.exe
BINARY_NAME=hw1-atkins
OUTPUT_NAME=z80output.txt
OUTPUT_EXISTS:=$(or $(and $(wildcard $(OUTPUT_NAME)),1),0)

all:
	$(GOBUILD) -o $(BINARY_NAME_WIN)
	$(GOBUILD) -o $(BINARY_NAME)

clean:
	$(GOCLEAN)

run:
	$(GORUN) main.go > $(OUTPUT_NAME)

dist: tar
	$(info "Making archive: $(ARCHIVE)")
	git archive -o $(ARCHIVE) HEAD^{tree}

tar:
	ARCHIVE=../$(BINARY_NAME).tar
	
zip:
	ARCHIVE=../$(BINARY_NAME).zip
	