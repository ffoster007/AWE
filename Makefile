.PHONY: build install uninstall clean run help

# ชื่อโปรแกรม
BINARY_NAME=awe

# ตำแหน่งที่จะติดตั้ง
INSTALL_PATH=$(shell go env GOPATH)/bin

# Default target
all: build

# Build โปรแกรม
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) ./cli
	@echo "Build complete!"

# ติดตั้งโปรแกรม
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@mkdir -p $(INSTALL_PATH)
	@cp $(BINARY_NAME) $(INSTALL_PATH)/
	@echo "Installation complete!"
	@echo "You can now run: $(BINARY_NAME) -h"

# ถอนการติดตั้ง
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Uninstall complete!"
