YELLOW = "\e[33m"
GREEN = "\e[32m"
RED = "\e[31m"

GENERATED_PROTO_CODE = pkg/api/nlo.pb.go
PROTO_FILE = api/nlo.proto




$(GENERATED_PROTO_CODE): $(PROTO_FILE)
	@printf $(YELLOW)
	protoc -I api --go_out=plugins=grpc:API $(PROTO_FILE)


.PHONY: build
build: $(GENERATED_PROTO_CODE) #db_up
	@printf $(GREEN)
	go build -v ./cmd/server


.PHONY: fclean
fclean:
	@printf $(RED)
	rm -f server




.DEFAULT_GOAL := build