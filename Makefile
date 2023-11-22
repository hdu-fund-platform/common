PROTO_FILES = $(shell  find ./protos -name *.proto)

.PHONY: gen
gen:
	protoc --go_out=.\
 		   --go-grpc_out=.\
 		   --grpc-gateway_out=. \
 		   --openapiv2_out=./openapi \
 		   $(PROTO_FILES)



.PHONY: init
init:
	go get google.golang.org/grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2

	@echo "done."