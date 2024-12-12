# Default target: Display help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run NUM=<number> LANG=<language>  # Run specified Go file"
	@echo "  make formatting                        # Format all Go files"

# Define variables for common commands and options
GOFMT = gofmt -w ./
GORUN = go run

# Target to format code
formatting:
	$(GOFMT)

# Pattern rule to run Go files based on directory and language
run:
	make formatting;
	$(GORUN) $(NUM)/$(NUM)_$(LANG).go
