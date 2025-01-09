generate: generate-auth generate-chat

generate-auth:
	protoc --go_out=services/auth --go_opt=paths=import \
	--go-grpc_out=services/auth --go-grpc_opt=paths=import \
	proto/auth/*.proto

generate-chat:
	protoc --go_out=services/chat --go_opt=paths=import \
	--go-grpc_out=services/chat --go-grpc_opt=paths=import \
	proto/chat/*.proto
