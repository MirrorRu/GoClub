GEARS_BIN:=$(CURDIR)/GEARS/bin
GEARS_DBDRIVER=postgres
GEARS_DBSTRING=postgres://postgres:postgres@localhost:5432/goclub?sslmode=disable
GEARS_MIGRATION_DIR=./GEARS/migrations

#starting stub
.PHONY: hello 
hello:
	$(info Hello!)
	@grep '^[^\.]\w.*\: ' Makefile

.PHONY: setup
setup: .bin-goose

.PHONY: .bin-goose
.bin-goose:
	$(info Installing goose binary...)
	GOBIN=$(GEARS_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: migrate-up
migrate-up: 
	$(GEARS_BIN)/goose -dir $(GEARS_MIGRATION_DIR) $(GEARS_DBDRIVER) $(GEARS_DBSTRING) up

.PHONY: migrate-down
migrate-down: 
	$(GEARS_BIN)/goose -dir $(GEARS_MIGRATION_DIR) $(GEARS_DBDRIVER) $(GEARS_DBSTRING) down

# .start:
# 	docker-compose up --build -d 

# build-all:
# 	cd cart && GOOS=linux GOARCH=amd64 make build
# 	cd GEARS && GOOS=linux GOARCH=amd64 make build
# 	cd notifier && GOOS=linux GOARCH=amd64 make build

# run-all: build-all .start .migrate

# # запуск окружения для проверки разработки
# run-develop:
# 	docker-compose --file docker-compose-develop.yml up

# .stop:
# 	docker-compose down
# #	docker rmi cart GEARS

# stop-all: .stop

# # остановка окружения проверки разработки
# stop-develop:
# 	docker-compose down

# docker-up:
# 	docker-compose up -d

# docker-down:
# 	docker-compose down
