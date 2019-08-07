UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

ARTICLE_ID:=1
ARTICLE_TITLE:=title
ARTICLE_BODY:=body

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-private:
	curl -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

req-article-root:
	@curl $(HOST):$(PORT)/articles

req-article-post:
	@curl -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'

req-article-update:
	@curl -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID) -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'

req-article-get:
	@curl $(HOST):$(PORT)/articles/$(ARTICLE_ID)

req-article-delete:
	@curl -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles/$(ARTICLE_ID)

req-public:
	curl $(HOST):$(PORT)/public

database-init:
	make -C ../database init

article-test-data:
	mysql -u root -h localhost --protocol tcp -e "use treasure_app; INSERT INTO treasure_app.article (title, body, ctime, utime) VALUES ('title', 'body', DEFAULT, DEFAULT)" -p
