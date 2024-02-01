# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.0] - 2024-02-01

### Added

- More tests for `DoAnnotation`

### Changed

- Bump Go to 1.18
- Refactor empty interfaces to `any`

## [1.1.1] - 2024-01-19

### Added

- Adding validation for file

### Changed

- Proper errors for `WordsFile`

## [1.1.0] - 2024-01-11

### Added

- Adding `WithSuffix` API
- Adding `Suffix` type in "Core" package
- Adding `DoAnnotation` API
- File state and emptiness checking for `WordsFile`
- Adding helper functions:
  - `GetBy`
  - `FindBy`

### Changed

- Proper error message for benchmarks

## [1.0.4] - 2023-12-21

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

[1.2.0]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.2.0
[1.1.1]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.1.1
[1.1.0]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.1.0
[1.0.4]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.4
[1.0.3]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.3
[1.0.2]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.2
[1.0.1]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.1
[1.0.0]: https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.0.0
