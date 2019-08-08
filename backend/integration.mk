UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

ARTICLE_ID:=1
ARTICLE_TITLE:=title
ARTICLE_BODY:=body

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-articles:
	curl -v $(HOST):$(PORT)/articles

req-img-pei:
	curl -v $(HOST):$(PORT)/img/pei.png

req-articles-get:
	curl -v $(HOST):$(PORT)/articles/$(ARTICLE_ID)

req-articles-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'

req-articles-update:
	curl -v -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID) -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'

req-articles-delete:
	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID)

req-public:
	curl -v $(HOST):$(PORT)/public

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

database-init:
	make -C ../database init
