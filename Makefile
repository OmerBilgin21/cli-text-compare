ENTRY_POINT = ./cmd/main.go
APP_NAME = clidiff
BIN_DIR  = ./bin

.PHONY: clean test

test:
	go test -v ./test/

all: clean build zip

# name:GOOS:GOARCH
PLATFORMS := \
    linux-amd64:linux:amd64 \
    linux-arm64:linux:arm64 \
    macos-amd64:darwin:amd64 \
    macos-arm64:darwin:arm64 \
    win-amd64:windows:amd64 \
    win-arm64:windows:arm64

name   = $(word 1,$(subst :, ,$(1)))
goos   = $(word 2,$(subst :, ,$(1)))
goarch = $(word 3,$(subst :, ,$(1)))
ext    = $(if $(filter windows,$(call goos,$(1))),.exe,)
bin    = $(BIN_DIR)/$(APP_NAME)-$(call name,$(1))$(call ext,$(1))
zipp   = $(BIN_DIR)/$(APP_NAME)-$(call name,$(1)).zip

define build_rule
$(call bin,$(1)): $(BIN_DIR)
	GOOS=$(call goos,$(1)) GOARCH=$(call goarch,$(1)) go build -o $$@ $(ENTRY_POINT)
endef

define zip_rule
$(call zipp,$(1)): $(call bin,$(1))
	zip -j $$@ $$<
endef

$(foreach p,$(PLATFORMS),$(eval $(call build_rule,$(p))))
$(foreach p,$(PLATFORMS),$(eval $(call zip_rule,$(p))))

BINS := $(foreach p,$(PLATFORMS),$(call bin,$(p)))
ZIPS := $(foreach p,$(PLATFORMS),$(call zipp,$(p)))

build: clean $(BINS) zip

zip: $(ZIPS)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)
