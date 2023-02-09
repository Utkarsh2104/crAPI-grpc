gen:
	protoc -I community_proto community_proto/*.proto --go_out=pb --go-grpc_out=pb