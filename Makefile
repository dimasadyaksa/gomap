GO ?= go
TPARSE ?= tparse

setup: install-pre-commit register-pre-commit install-tparse install-gofumpt install-errcheck install-golangci-lint install-staticcheck

test:
	$(GO) test -race -coverprofile=coverage.out -timeout 30s -json ./... | $(TPARSE) -all

install-tparse: # install tparse if not installed yet.
	@if ! command -v tparse > /dev/null; then \
		echo "Installing tparse"; \
		$(GO) install github.com/mfridman/tparse@latest; \
	fi

install-pre-commit:
	@if ! command -v pre-commit > /dev/null; then \
		echo "Installing pre-commit"; \
		pip3 install pre-commit; \
	fi

install-gofumpt:
	@if ! command -v gofumpt > /dev/null; then \
		echo "Installing gofumpt"; \
		$(GO) install mvdan.cc/gofumpt@latest; \
	fi

install-staticcheck:
	@if ! command -v staticcheck > /dev/null; then \
		echo "Installing staticcheck"; \
		$(GO) install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi

install-golangci-lint:
	@if ! command -v golangci-lint > /dev/null; then \
		echo "Installing golangci-lint"; \
		$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi

install-errcheck:
	@if ! command -v errcheck > /dev/null; then \
		echo "Installing errcheck"; \
		$(GO) install github.com/kisielk/errcheck@latest; \
	fi

register-pre-commit:
	@pre-commit install --hook-type commit-msg