# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Adding get and find with prefix

## [1.0.4] - 2023-12-

### Added

- Adding "misspell" to Makefile
- Adding new badges to "README.md"

### Fixed

- Package import in "example_test.go"

## [1.0.3] - 2023-12-17

### Added

- File validations for `WordsFile`
- Separate github workflow files for test and coverage

### Changed

- Refine tests

## [1.0.2] - 2023-12-15

### Added

- Test coverage to Makefile
- [Go Report Card](https://goreportcard.com/badge) and Coverage (https://img.shields.io) badges to README
- Checking for "Scanner" errors in "WordsFile.FindUnsafe"
- Tests for "Instantiation"
- More tests for "WordsFile"

### Changed

- Rename "Collection" function in "internal" to "Treasure"
- Some refinement in "internal_test.go"

## [1.0.1] - 2023-12-14

### Fixed

- misspelled English words

## [1.0.0] - 2023-12-14

### Added

- WordsRepository
- WordsCollection
- WordsFile

[unreleased]: https://github.com/saleh-rahimzadeh/go-words/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.0
[1.0.1]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.1
[1.0.2]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.2
[1.0.3]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.3
