YELLOW = "\e[33m"
GREEN = "\e[32m"
RED = "\e[31m"

GENERATED_PROTO_CODE = pkg/api/nlo.pb.go
PROTO_FILE = api/nlo.proto




$(GENERATED_PROTO_CODE): $(PROTO_FILE)
	@printf $(YELLOW)
	protoc -I api --go_out=plugins=grpc:API $(PROTO_FILE)

.PHONY: build
build: $(GENERATED_PROTO_CODE) db_up
	@printf $(GREEN)
	go build -v ./cmd/server
	go build -v ./cmd/client

.PHONY: db_up
db_up:
	@printf $(GREEN)
	docker run --rm --name my_postgres \
	-e POSTGRES_PASSWORD=qwerty1234 \
	-e POSTGRES_USER=kshanti \
	-e POSTGRES_DB=anomaly_db \
	-p 5432:5432 -d postgres

.PHONY: db_down
db_down:
	docker stop my_postgres

.PHONY: fclean
fclean: db_down
	@printf $(RED)
	rm -f server
	rm -f client

.PHONY: re
re: fclean build


.DEFAULT_GOAL := build