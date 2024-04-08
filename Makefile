BINARY_NAME=MarkDown.app
APP_NAME=MarkDown
VERSION=1.0.0

build:
	rm -rf $(BINARY_NAME)
	rm -f fyne-md
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release

run:
	go run .

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -rf $(BINARY_NAME)
	@echo "Cleaned up"

test:
	go test -v ./...
