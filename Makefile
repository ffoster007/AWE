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
	go build -o $(BINARY_NAME) .
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

# ลบไฟล์ที่ build
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Clean complete!"

# รันโปรแกรมโดยตรง (สำหรับทดสอบ)
run: build
	@./$(BINARY_NAME)

# แสดงคำสั่งที่ใช้ได้
help:
	@echo "Available targets:"
	@echo "  make build      - Build the project"
	@echo "  make install    - Build and install to GOPATH/bin"
	@echo "  make uninstall  - Remove installed binary"
	@echo "  make clean      - Remove built binary"
	@echo "  make run        - Build and run locally"
	@echo "  make help       - Show this help message"