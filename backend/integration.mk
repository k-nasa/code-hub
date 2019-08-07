UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-private:
	curl -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

req-article-root:
	@curl $(HOST):$(PORT)/articles

req-article-post:
	@curl -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/articles -d '{"title": "hoge", "body": "fuga"}'

ID:=1
req-article-get:
	@curl $(HOST):$(PORT)/articles/$(ID)

req-public:
	curl $(HOST):$(PORT)/public

database-init:
	make -C ../database init

article-test-data:
	mysql -u root -h localhost --protocol tcp -e "use treasure_app; INSERT INTO treasure_app.article (title, body, ctime, utime) VALUES ('title', 'body', DEFAULT, DEFAULT)" -p
