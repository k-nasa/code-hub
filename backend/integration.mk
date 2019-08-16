export GO111MODULE := on

UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

CODE_ID :=1
CODE_TITLE:=俺が考えた最強のコード
CODE_BODY:=hogehoge
CODE_STATUS:=public

ARTICLE_COMMENT_BODY:=bodycomment

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-public:
	curl -v $(HOST):$(PORT)/public

req-code-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/codes -d '{"title": "$(CODE_TITLE)", "body": "$(CODE_BODY)", "status": "$(CODE_STATUS)"}'

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

database-init:
	make -C ../database init