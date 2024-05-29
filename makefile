build:
	@go build -o app main.go

run: build
	@go run main.go

clean:
	@go clean
	@del app

gen:
	@protoc --go_out=pb --go-grpc_out=pb proto/*.proto