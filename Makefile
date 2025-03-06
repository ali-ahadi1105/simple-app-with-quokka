BINARY_NAME=quokkaApp

build:
	@go mod vendor
	@echo "Building Quokka..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Quokka built!"

run: build
	@echo "Starting Quokka..."
	@./tmp/${BINARY_NAME} &
	@echo "Quokka started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Quokka..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Quokka!"

restart: stop start