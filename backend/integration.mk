UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-private:
	curl -H "Authorization: Bearer $(shell cat ./.idToken)" $(HOST):$(PORT)/private

req-public:
	curl $(HOST):$(PORT)/public

database-init:
	make -C ../database init
