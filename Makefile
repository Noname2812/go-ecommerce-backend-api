# include .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

GOOSE_DBSTRING ?= $(STR_MYSQL)
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

PROTO_SRC_DIR := proto
PROTO_OUT_BASE := internal/common/protogen

# Tên của ứng dụng của bạn
APP_NAME := server

# Detect OS
ifeq ($(OS),Windows_NT)
    detected_OS := Windows
    SET_CMD := set
    AND_CMD := &&
else
    detected_OS := $(shell uname -s)
    SET_CMD := export
    AND_CMD := ;
endif

# test
all:
	@echo "The message is: $(MESSAGE)"
	@echo "The count is: $(COUNT)"

print_vars:
	@echo "MESSAGE (from make): $(MESSAGE)"
	@echo "COUNT (from make): $(COUNT)"
	@echo "STR_MYSQL: $(STR_MYSQL)"
	@echo "GOOSE_DBSTRING: $(GOOSE_DBSTRING)"

# Chạy ứng dụng
docker_build:
	docker-compose -f environment/docker-compose-dev.yml up -d --build
	docker-compose ps

docker_down:
	docker-compose -f environment/docker-compose-dev.yml down
	docker-compose -f environment/kafka/docker-compose-kafka-single.yml down

dev:
	go run ./cmd/$(APP_NAME)

docker_up:
	docker-compose -f environment/kafka/docker-compose-kafka-single.yml up -d
	docker-compose -f environment/docker-compose-dev.yml up -d

# Migration commands - Windows compatible
up_by_one:
ifeq ($(detected_OS),Windows)
	$(SET_CMD) "GOOSE_DRIVER=$(GOOSE_DRIVER)" $(AND_CMD) $(SET_CMD) "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" $(AND_CMD) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
else
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
endif

# create new a migration
create_migration:
	goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
ifeq ($(detected_OS),Windows)
	$(SET_CMD) "GOOSE_DRIVER=$(GOOSE_DRIVER)" $(AND_CMD) $(SET_CMD) "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" $(AND_CMD) goose -dir=$(GOOSE_MIGRATION_DIR) up
else
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
endif

downse:
ifeq ($(detected_OS),Windows)
	$(SET_CMD) "GOOSE_DRIVER=$(GOOSE_DRIVER)" $(AND_CMD) $(SET_CMD) "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" $(AND_CMD) goose -dir=$(GOOSE_MIGRATION_DIR) down
else
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
endif

resetse:
ifeq ($(detected_OS),Windows)
	$(SET_CMD) "GOOSE_DRIVER=$(GOOSE_DRIVER)" $(AND_CMD) $(SET_CMD) "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" $(AND_CMD) goose -dir=$(GOOSE_MIGRATION_DIR) reset
else
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset
endif

sqlgen:
	sqlc generate

# protoc run as gitbash on windows
proto:
	@echo "🔧 Generating all .proto files..."
	@for file in $(PROTO_SRC_DIR)/*.proto; do \
		filename=$$(basename $$file .proto); \
		outdir=$(PROTO_OUT_BASE)/$$filename; \
		mkdir -p $$outdir; \
		echo "🛠  Generating $$file → $$outdir"; \
		protoc \
			--proto_path=$(PROTO_SRC_DIR) \
			--go_out=$$outdir --go_opt=paths=source_relative \
			--go-grpc_out=$$outdir --go-grpc_opt=paths=source_relative \
			$$file; \
	done
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up swag air print_vars proto