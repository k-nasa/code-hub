export GO111MODULE := on

UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

CODE_ID :=1
USER_ID :=1
CODE_TITLE:=俺が考え強のコード9
CODE_BODY:=unc AllCodesWithUsedb *sqlx.DB) ([]model.CodeWithUser, error) {
CODE_STATUS:=public

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

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

database-init:
	make -C ../database init
