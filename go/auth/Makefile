# Detect OS
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
	OS := macOS
else ifeq ($(UNAME_S),Linux)
	IS_UBUNTU := $(shell lsb_release -is 2>/dev/null)
	ifeq ($(IS_UBUNTU),Ubuntu)
		OS := Ubuntu
	else
		OS := Unsupported
	endif
else
	OS := Unsupported
endif

# Makefile targets
.PHONY: check-os
check-os:
	@echo "💡 Detected OS: $(OS)"
	@if [ "$(OS)" = "Unsupported" ]; then \
		echo "❌ Error: Unsupported OS. This Makefile only supports macOS and Ubuntu." >&2; \
		exit 1; \
	fi
	@echo "✅ Proceeding on $(OS)..."

.PHONY: setup
setup:
	@echo "🔧 Setting up environment..."
	@if [ "$(OS)" = "macOS" ]; then \
		brew install protobuf; \
	elif [ "$(OS)" = "Ubuntu" ]; then \
		sudo apt update; \
		sudo apt install -y protobuf-compiler; \
	fi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@echo "✅ Environment setup complete."

.PHONY: generate
generate:
	@echo "✨ Generating gRPC code from .proto files..."
	protoc -I ./proto \
		--go_out ./proto --go_opt paths=source_relative \
		--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
		proto/services.proto
	@echo "✅ Code generation complete."
