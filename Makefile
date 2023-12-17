#╔═════════════════════════════════════════════════════════════════════════════════════════════════╗
#║ Copyright (C) 2023 Saleh Rahimzadeh                                                             ║
#║ https://github.com/saleh-rahimzadeh/go-words                                                    ║
#╚═════════════════════════════════════════════════════════════════════════════════════════════════╝

.DEFAULT_GOAL := help

# ------------------------------------------------------------------------------

COVERAGE_FILE          = testdata/coverage.out
COVERAGE_ANALYSIS_FILE = testdata/coverage.analysis.out
BRANCH := v1.0

# ------------------------------------------------------------------------------

help:
	@egrep "^##" Makefile|sed 's/##//g'

release:
	@echo "► release"
	@test -n "$(BRANCH)" || (echo "Error: BRANCH arg is empty" ; exit 1)
	@test -n "$(TAG)"    || (echo "Error: TAG arg is empty" ; exit 1)
	@echo "Branch: $(BRANCH)"
	@echo "Tag:    $(TAG)"
	git checkout main
	git merge --no-ff $(BRANCH)
	git tag '$(TAG)'
.PHONY:release

## fmt           : Applies standard formatting (whitespace, indentation, ...).
fmt:
	@echo "► fmt"
	go fmt ./...
.PHONY:fmt

## lint          : Using static analysis, it finds bugs and performance issues, and enforces style rules.
lint:
	@echo "► lint"
	staticcheck ./...
.PHONY:lint

## vet           : Find subtle errors and issues where not caught by the compilers and code may not work as intended.
vet:
	@echo "► vet"
	go vet ./...
.PHONY:vet

## analyze       : Analyze code using : ► vet ► lint ► fmt
analyze: vet lint fmt
.PHONY:analyze

## coverage      : Create test coverage file
coverage:
	@echo "► coverage"
	go test ./... -covermode=count -coverprofile=$(COVERAGE_FILE)
	go tool cover -func=$(COVERAGE_FILE) -o=$(COVERAGE_ANALYSIS_FILE)
	@echo -n -e '\n= '
	@tail -n 1 $(COVERAGE_ANALYSIS_FILE) | sed 's/total://g;s/	*//g;s/(statements)//g'
.PHONY:coverage

## coverage_html : Create HTML representation of coverage file
coverage_html:
	@echo "► coverage_html"
	go test ./... -covermode=count -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE)
.PHONY:coverage_html

## misspell      : Check commonly misspelled English words in source files by "github.com/client9/misspell"
misspell:
	@echo "► misspell"
	misspell .
.PHONY:misspell

## generate      : Generate fake large words for benchmarking
generate:
	@cd testdata/scripts; \
	./generate_words.sh
.PHONY:generate

## generate_ps   : Generate fake large words for benchmarking by PowerShell script
generate_ps:
	@cd testdata/scripts; \
	pwsh -File ./generate_words.ps1
.PHONY:generate_ps
