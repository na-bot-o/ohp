#parameter
#references https://frasco.io/golang-dont-afraid-of-makefiles-785f3ec7eb32
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ohp
DATA_FILE=~/.ohp


all: reset test build

build:
	$(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST) -v ./...
setup:
	touch $(DATA_FILE)
#recreate .ohp file if have this
reset:
	@if test -f $(DATA_FILE); \
		then rm -f $(DATA_FILE); \
	fi
	$(MAKE) setup
