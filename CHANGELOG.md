# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic Versioning].

## [Unreleased]

- Will list future updates here

## [1.0.2] - 2023-05-27

### Added

- Added `TODO.md` file
- Added `build.yml` to `.github/workflows`
- Added `lint.yml` to `.github/workflows`

### Changed

- Renamed `market.go` to `go-nyse-time.go`
- Moved `holiday.go` to `src` directory
- Moved `hours.go` to `src` directory
- Moved `go-nys-time.go` to `src` directory
- Updated `.gitignore` file

### Removed

- Removed `go.yml` from `.github/workflows`

## [1.0.1] - 2023-05-23

### Changed

- Updated files to reflect new name

## [1.0.0] - 2023-01-18

### Added

- Better Commenting

### Changed

- Updated `CHANGELOG.md` to include changes

## [0.0.3] - 2023-01-18

### Changed

- Changed repo title to `go-nyse-time`
- Updated `README.md` to include usage examples
- Updated `CHANGELOG.md` to include changes

## [0.0.2] - 2023-01-18

### Changed

- Changed package name from `nyse` to `nyse_time`
- Updated `README.md` to include usage examples
- Updated `CHANGELOG.md` to include changes

### Removed

- Large amount of public functions that needed to be private

## [0.0.1] - 2023-01-17

### Added

- `IsMarketOpen()` - Checks if the market is open at the current time
  - Returns `true` if the market is open
- `IsMarketOpenCustom(time.Time)` - Checks if the market is open at a specific time
  - Returns `true` if the market is open
- `IsHoliday(time.Time)` - Checks if a specific date is a holiday
  - Returns `true` if the date is a holiday
- `IsEarlyClose(time.Time)` - Checks if a specific date is an early close
  - Returns `true` if the date is an early close
- `IsWeekend(time.Time)` - Checks if a specific date is a weekend
  - Returns `true` if the date is a weekend

### Changed

- Updated `README.md` to include usage examples
- Updated `CHANGELOG.md` to include changes

<!-- Links -->
[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

<!-- Versions -->
[unreleased]: https://github.com/ggoodwin/go-nyse-time/compare/v1.0.2...DEVELOP
[1.0.2]: https://github.com/ggoodwin/go-nyse-time/compare/v1.0.0...v1.0.1
[1.0.1]: https://github.com/ggoodwin/go-nyse-time/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/ggoodwin/go-nyse-time/compare/v0.0.3...v1.0.0
[0.0.3]: https://github.com/ggoodwin/go-nyse-time/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/ggoodwin/go-nyse-time/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/ggoodwin/go-nyse-time/releases/tag/v0.0.1
