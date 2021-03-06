export GO111MODULE := on

UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

CODE_ID :=27
USER_ID :=27
CODE_TITLE:=俺が考え強のコード9
COMMENT_BODY:=すごいね、、、
CODE_BODY:=unc AllCodesWithUsedb *sqlx.DB) ([]model.CodeWithUser, error) {
CODE_STATUS:=public
LANGUAGE:=golang

ARTICLE_COMMENT_BODY:=bodycomment

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-public:
	curl -v $(HOST):$(PORT)/public

req-code-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/codes -d '{"title": "$(CODE_TITLE)", "body": "$(CODE_BODY)", "status": "$(CODE_STATUS)"}'

req-codes:
	curl -v $(HOST):$(PORT)/codes

req-code-get:
	curl -v $(HOST):$(PORT)/codes/$(CODE_ID)

req-codes-user:
	curl -v $(HOST):$(PORT)/users/codes

req-users-code:
	curl -v $(HOST):$(PORT)/users/$(USER_ID)/codes

req-code-by-title:
	curl -v $(HOST):$(PORT)/nasa/$(CODE_TITLE)

req-code-comments:
	curl -v $(HOST):$(PORT)/codes/$(CODE_ID)/comments

req-comment-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/comments -d '{"code_id": $(CODE_ID), "body": "$(COMMENT_BODY)" }'

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private


req-compile:
	curl -v -XPOST $(HOST):$(PORT)/compile -d '{"language": "$(LANGUAGE)", "body": "package main"}'

database-init:
	make -C ../database init
