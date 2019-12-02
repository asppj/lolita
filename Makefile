
protoDir=./api/proto/
task=task
plan=plan
rpc:
	protoc --go_out=plugins=grpc:. ${protoDir}${task}/*.proto
	protoc --go_out=plugins=grpc:. ${protoDir}${plan}/*.proto