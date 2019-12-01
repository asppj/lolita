
protoDir=./proto/
task=task
rpc:
	protoc --go_out=plugins=grpc:. ${protoDir}${task}/*.proto