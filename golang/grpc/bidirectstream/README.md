protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/bidirectstream.proto

go run server/server.go

grpcui -plaintext 0.0.0.0:50058