

protoDir=./api/proto/
task=task
plan=plan
rpc:
	protoc --go_out=plugins=grpc:. ${protoDir}${task}/*.proto
	protoc --go_out=plugins=grpc:. ${protoDir}${plan}/*.proto

build:
	go get -v
	go test ./test
	go build -v


publish:
	docker build . -t "t-go-opentrace:latest"
	docker push asppj/t-go-opentrace:latest

compose:
	docker-compose -f docker-compose.yml down
	docker-compose -f docker-compose.yml up -d --force-recreate

start:
	go run main.go