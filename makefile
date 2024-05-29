build:
	@go build -o app main.go

run:
	@go run main.go

clean:
	@go clean
	@rm app
