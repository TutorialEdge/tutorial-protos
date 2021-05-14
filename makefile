build:
	protoc --go_out=plugins=grpc:. protos/**/*.proto