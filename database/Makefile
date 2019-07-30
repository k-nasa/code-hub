DBNAME:=treasure_app
DATASOURCE:=root:password@tcp(127.0.0.1:3306)/$(DBNAME)

install:
	go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up

docker-compose/down:
	docker-compose down

mysql:
	mysql -u root -h localhost --protocol tcp -p

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/status:
	goose -dir migrations mysql "$(DATASOURCE)" status

migrate/up:
	goose -dir migrations mysql "$(DATASOURCE)" up

__migrate/down:
	goose -dir migrations mysql "$(DATASOURCE)" down
