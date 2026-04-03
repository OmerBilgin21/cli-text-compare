APP_NAME = textdiff
VERSION  = v2.0.0
BIN_DIR  = ./bin

.PHONY: clean test

test:
	go test ./test/ -v

all: clean build zip

build: $(BIN_DIR)/$(APP_NAME)-linux-amd64 \
       $(BIN_DIR)/$(APP_NAME)-linux-arm64 \
       $(BIN_DIR)/$(APP_NAME)-macos-amd64 \
       $(BIN_DIR)/$(APP_NAME)-macos-arm64 \
       $(BIN_DIR)/$(APP_NAME)-win-amd64.exe \
       $(BIN_DIR)/$(APP_NAME)-win-arm64.exe

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(BIN_DIR)/$(APP_NAME)-linux-amd64: $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-linux-arm64: $(BIN_DIR)
	GOOS=linux GOARCH=arm64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-macos-amd64: $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-macos-arm64: $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-win-amd64.exe: $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-win-arm64.exe: $(BIN_DIR)
	GOOS=windows GOARCH=arm64 go build -o $@ ./cmd/main.go

zip: $(BIN_DIR)/$(APP_NAME)-linux-amd64.zip \
     $(BIN_DIR)/$(APP_NAME)-linux-arm64.zip \
     $(BIN_DIR)/$(APP_NAME)-macos-amd64.zip \
     $(BIN_DIR)/$(APP_NAME)-macos-arm64.zip \
     $(BIN_DIR)/$(APP_NAME)-win-amd64.zip \
     $(BIN_DIR)/$(APP_NAME)-win-arm64.zip

$(BIN_DIR)/$(APP_NAME)-linux-amd64.zip: $(BIN_DIR)/$(APP_NAME)-linux-amd64
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-linux-arm64.zip: $(BIN_DIR)/$(APP_NAME)-linux-arm64
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-macos-amd64.zip: $(BIN_DIR)/$(APP_NAME)-macos-amd64
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-macos-arm64.zip: $(BIN_DIR)/$(APP_NAME)-macos-arm64
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-win-amd64.zip: $(BIN_DIR)/$(APP_NAME)-win-amd64.exe
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-win-arm64.zip: $(BIN_DIR)/$(APP_NAME)-win-arm64.exe
	zip -j $@ $<

clean:
	rm -rf $(BIN_DIR)
