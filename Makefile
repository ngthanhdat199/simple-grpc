run-api:
	@echo "Running the application"
	@go run main.go

build-api:	
	docker build -t root-server -f ./Dockerfile .
