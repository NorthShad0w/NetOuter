LD_FLAGS := -s -w
BIN_DIR := target/release
BIN_NAME := NetOuter

default: clean darwin linux windows integrity

clean:
	$(RM) $(BIN_DIR)/$(BIN_NAME)*
	go clean -x

install:
	go install

darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-darwin-amd64' ./cmd/netouter/main.go

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-linux-amd64' ./cmd/netouter/main.go

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-windows-amd64.exe' ./cmd/netouter/main.go

integrity:
	cd $(BIN_DIR) && shasum *
