APP_NAME=machine-auth
SRC_GEN_FINGERPRINT=cmd/gen-fingerprint.go
SRC_GEN_LICENSE=cmd/gen-license.go
SRC_VERIFY=cmd/verify.go
OUT_DIR=bin

# 只构建 Linux amd64
PLATFORM=linux/amd64

all: build

build:
	mkdir -p $(OUT_DIR)
	# 打包 gen-fingerprint.go
	GOOS=linux GOARCH=amd64 \
		out=$(OUT_DIR)/gen-fingerprint-linux-amd64; \
		echo "🔧 Building $$out..."; \
		GOOS=linux GOARCH=amd64 go build -o $$out $(SRC_GEN_FINGERPRINT)

	# 打包 gen-license.go
	GOOS=linux GOARCH=amd64 \
		out=$(OUT_DIR)/gen-license-linux-amd64; \
		echo "🔧 Building $$out..."; \
		GOOS=linux GOARCH=amd64 go build -o $$out $(SRC_GEN_LICENSE)

	# 打包 verify.go
	GOOS=linux GOARCH=amd64 \
		out=$(OUT_DIR)/verify-linux-amd64; \
		echo "🔧 Building $$out..."; \
		GOOS=linux GOARCH=amd64 go build -o $$out $(SRC_VERIFY)

clean:
	rm -rf $(OUT_DIR)

.PHONY: all build clean
