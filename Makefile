GO_MODULE := github.com/MuhAndriansyah/grpc-bank-project

.PHONY: protoc-go
protoc-go:
	# root of all proto relative imports
	protoc \
		--proto_path=proto \
		--go_opt=module=${GO_MODULE} --go_out=. \
		--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
		./proto/bank/v1/*.proto ./proto/bank/v1/type/*.proto \
