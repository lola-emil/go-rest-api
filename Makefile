APP_NAME=go-rest-api
BUILD_DIR=bin
MAIN=./cmd/$(APP_NAME)

dev:
	go run ./cmd/$(APP_NAME)/main.go


build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN)

run: build
	./$(BUILD_DIR)/$(APP_NAME)


clean:
	rm -rf $(BUILD_DIR)