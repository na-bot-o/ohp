#parameter
#references https://frasco.io/golang-dont-afraid-of-makefiles-785f3ec7eb32
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=ohp
DATA_FILE=~/.ohp


all: init reset test install

init:
	$(GOGET) github.com/golang/dep/cmd/dep
	dep ensure
build:
	$(GOBUILD) -o $(BINARY_NAME)
install:
	$(GOINSTALL)
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
