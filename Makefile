APP_NAME = textdiff
VERSION  = v1.0.0
BIN_DIR  = bin

.PHONY: all clean

all: build zip

# ----- Build section -----

build: $(BIN_DIR)/$(APP_NAME)-linux \
       $(BIN_DIR)/$(APP_NAME)-macos-arm64 \
       $(BIN_DIR)/$(APP_NAME)-win.exe

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(BIN_DIR)/$(APP_NAME)-linux: $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-macos-arm64: $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $@ ./cmd/main.go

$(BIN_DIR)/$(APP_NAME)-win.exe: $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $@ ./cmd/main.go

# ----- Zip section -----

zip: $(BIN_DIR)/$(APP_NAME)-linux.zip \
     $(BIN_DIR)/$(APP_NAME)-macos-arm64.zip \
     $(BIN_DIR)/$(APP_NAME)-win.zip

$(BIN_DIR)/$(APP_NAME)-linux.zip: $(BIN_DIR)/$(APP_NAME)-linux
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-macos-arm64.zip: $(BIN_DIR)/$(APP_NAME)-macos-arm64
	zip -j $@ $<

$(BIN_DIR)/$(APP_NAME)-win.zip: $(BIN_DIR)/$(APP_NAME)-win.exe
	zip -j $@ $<

clean:
	rm -rf $(BIN_DIR)
