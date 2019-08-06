export GO111MODULE := on

GOOS := linux
GOARCH := amd64
UID := demo
ID_TOKEN := $(shell cat ./.idToken)
PORT := 1991
HOST := localhost
AUTH_HEADER := -H "Authorization: Bearer $(ID_TOKEN)"

init: .env

.env:
	cp $@.example $@

go:
	gofmt -s -w .

run:
	go run cmd/api/main.go

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build cmd/api/main.go

create-token:
	go run ./cmd/customtoken/main.go $(UID)

req-private:
	curl $(AUTH_HEADER) $(HOST):$(PORT)/private

req-public:
	curl $(HOST):$(PORT)/public

database-init:
	make -C ../database init
