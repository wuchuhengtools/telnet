# Define the binary name and the output format
BINARY_NAME := app
OUTPUT_DIR := assets
# Get the version from the cli, like: `make build VERSION=1.0.0`
VERSION := $(if $(VERSION),$(VERSION),0.1.0)

# Define the build targets
PLATFORMS := darwin_amd64 darwin_arm64 linux_amd64 windows_amd64

# Function to build for a specific platform
define build_platform
	@echo "Building for $(1)..."
	GOOS=$(word 1, $(subst _, ,$(1))) GOARCH=$(word 2, $(subst _, ,$(1))) go build -o $(OUTPUT_DIR)/$(BINARY_NAME)_$(1)_$(VERSION) main.go
endef

# Build all targets
build: $(PLATFORMS:%=$(OUTPUT_DIR)/$(BINARY_NAME)_%)

# Build for each platform
$(OUTPUT_DIR)/$(BINARY_NAME)_darwin_amd64:
	$(call build_platform,darwin_amd64)

$(OUTPUT_DIR)/$(BINARY_NAME)_darwin_arm64:
	$(call build_platform,darwin_arm64)

$(OUTPUT_DIR)/$(BINARY_NAME)_linux_amd64:
	$(call build_platform,linux_amd64)

$(OUTPUT_DIR)/$(BINARY_NAME)_windows_amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_DIR)/$(BINARY_NAME)_windows_amd64_$(VERSION).exe

# Ensure the assets directory exists
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Clean up built files
clean:
	rm -f $(OUTPUT_DIR)/*