#╔═════════════════════════════════════════════════════════════════════════════════════════════════╗
#║ Copyright (C) 2023 Saleh Rahimzadeh                                                             ║
#║ https://github.com/saleh-rahimzadeh/go-words                                                    ║
#╚═════════════════════════════════════════════════════════════════════════════════════════════════╝

.DEFAULT_GOAL := help

# ------------------------------------------------------------------------------

COVERAGE_FILE = testdata/coverage.out

# ------------------------------------------------------------------------------

help:
	@egrep "^##" Makefile|sed 's/##//g'

## fmt      : Applies standard formatting (whitespace, indentation, ...).
fmt:
	@echo "► fmt"
	go fmt ./...
.PHONY:fmt

## lint     : Using static analysis, it finds bugs and performance issues, and enforces style rules.
lint:
	@echo "► lint"
	staticcheck ./...
.PHONY:lint

## vet      : Find subtle errors and issues where not caught by the compilers and code may not work as intended.
vet:
	@echo "► vet"
	go vet ./...
.PHONY:vet

## analyze  : Analyze code using : ► vet ► lint ► fmt
analyze: vet lint fmt
.PHONY:analyze

## coverage : Create test coverage file
coverage:
	@echo "► coverage"
	go test ./... -covermode=count -coverprofile=$(COVERAGE_FILE)
	go tool cover -func=$(COVERAGE_FILE) -o=$(COVERAGE_FILE)
	@echo -n -e '\n= '
	@tail -n 1 $(COVERAGE_FILE) | sed 's/total://g;s/	*//g;s/(statements)//g'
.PHONY:coverage

## generate : Generate fake large words for benchmarking
generate:
	@cd testdata/scripts; \
	./generate_words.sh
.PHONY:generate
