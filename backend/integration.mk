UID := demo
ID_TOKEN := $(shell cat ./.idToken)
PORT := 1991
HOST := localhost
AUTH_HEADER := -H "Authorization: Bearer $(ID_TOKEN)"

create-token:
	go run ./cmd/customtoken/main.go $(UID)

req-private:
	curl $(AUTH_HEADER) $(HOST):$(PORT)/private

req-public:
	curl $(HOST):$(PORT)/public

database-init:
	make -C ../database init
