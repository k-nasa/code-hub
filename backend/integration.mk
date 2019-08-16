export GO111MODULE := on

UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

ARTICLE_ID:=1
ARTICLE_TITLE:=title
ARTICLE_BODY:=body

ARTICLE_COMMENT_BODY:=bodycomment

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-public:
	curl -v $(HOST):$(PORT)/public

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

database-init:
	make -C ../database init
