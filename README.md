<div align="center">
	<h1><img alt="Time logo" src="https://github.com/ggoodwin/go-nyse-time/blob/master/timetable.png" height="300" /><br />
		New York Stock Exchange (NYSE) Time Go Library
	</h1>

[![GMan#0001](https://dcbadge.vercel.app/api/shield/179795086543028224)](https://discord.id/?prefill=179795086543028224) [![Go Package Devs](https://dcbadge.vercel.app/api/server/jwEYR2Dume)](https://discord.gg/jwEYR2Dume)

[![Go Reference](https://pkg.go.dev/badge/ggoodwin/go-nyse-time.svg)](https://pkg.go.dev/github.com/ggoodwin/go-nyse-time) [![Go Version](https://img.shields.io/github/go-mod/go-version/ggoodwin/go-nyse-time)](https://go.dev/) ![Size](https://img.shields.io/github/languages/code-size/ggoodwin/go-nyse-time) [![Last Commit](https://img.shields.io/github/last-commit/ggoodwin/go-nyse-time)](https://github.com/ggoodwin/go-nyse-time/commits/master) [![License](https://img.shields.io/github/license/ggoodwin/go-nyse-time)](https://github.com/ggoodwin/go-nyse-time/blob/master/LICENSE.md)

[![GoReportCard](https://goreportcard.com/badge/github.com/ggoodwin/go-nyse-time)](https://goreportcard.com/report/github.com/ggoodwin/go-nyse-time) [![CodeFactor](https://www.codefactor.io/repository/github/ggoodwin/go-nyse-time/badge)](https://www.codefactor.io/repository/github/ggoodwin/go-nyse-time) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/17f51d3e54264211b19220ce470783ae)](https://app.codacy.com/gh/ggoodwin/go-nyse-time/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![codecov](https://codecov.io/gh/ggoodwin/go-nyse-time/branch/master/graph/badge.svg?token=YNDB8EF3ZN)](https://codecov.io/gh/ggoodwin/go-nyse-time)

[![Build](https://github.com/ggoodwin/go-nyse-time/actions/workflows/build.yml/badge.svg)](https://github.com/ggoodwin/go-nyse-time/actions/workflows/build.yml) [![lint](https://github.com/ggoodwin/go-nyse-time/actions/workflows/lint.yml/badge.svg)](https://github.com/ggoodwin/go-nyse-time/actions/workflows/lint.yml) [![CodeQL](https://github.com/ggoodwin/go-nyse-time/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/ggoodwin/go-nyse-time/actions/workflows/github-code-scanning/codeql)

</div>
<hr/>

## 🌟 How it works

The library uses the current time to determine if the market is open, closed, or closed early. It also determines if the current day is a holiday.

## 📦 Installation and Usage

### Go

Make sure you have `Go` installed on your machine.

You can check by running the following command in the `console`

```plain
go version
```

If you don't have `Go` installed, you can download it from [here](https://go.dev/dl/).

### Add to your project

Run the following command in the `console`, in the `project directory`, to install the library with `go get`

```plain
go get github.com/ggoodwin/go-nyse-time
```

### Importing

Add the import to your `.go` file

```go
import nyse_time "github.com/ggoodwin/go-nyse-time"
```

## 💰 Usage

### Check if the Market is Open at Current Time

```go
// Returns: bool
open := nyse_time.IsMarketOpen()
```

### Check if the Market will be Open based on a specific time and date

```go
// Parameters: time.Time
// Returns: bool
open := nyse_time.IsMarketOpenCustom(time.Time)
```

### Is It A Holiday?

```go
// Parameters: time.Time
// Returns: bool
isHoliday := nyse_time.IsHoliday(time.Time)
```

### Is It An Early Close Day?

```go
// Parameters: time.Time
// Returns: bool
isEarlyClose := nyse_time.IsEarlyClose(time.Time)
```

### Is It A Weekend?

```go
// Parameters: time.Time
// Returns: bool
isWeekend := nyse_time.IsWeekend(time.Time)
```

## 💻 Dependencies

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report the issue using our [GITHUB-ISSUES] or check out the [SECURITY] details if it is security related.

If you'd like, I welcome any contributions. Please read the [CONTRIBUTING] document then [FORK] this library and submit a [PULL-REQUEST]. Make sure to click `compare across forks` to see your fork.

## ⚖️ License

This project is under the MIT License. See the [LICENSE] file for the full license text.

## 📜 Changes

Check out our [CHANGELOG]

## 👍🏻 Code of Conduct

Please read my [CODE-OF-CONDUCT] before contributing or engaging in discussions.

<!-- Links -->
[LICENSE]: https://github.com/ggoodwin/go-nyse-time/blob/master/LICENSE.md
[CHANGELOG]: https://github.com/ggoodwin/go-nyse-time/blob/master/CHANGELOG.md
[SECURITY]: https://github.com/ggoodwin/go-nyse-time/blob/master/SECURITY.md
[FORK]: https://github.com/ggoodwin/go-nyse-time/fork
[PULL-REQUEST]: https://github.com/ggoodwin/go-nyse-time/compare
[CODE-OF-CONDUCT]: https://github.com/ggoodwin/go-nyse-time/blob/master/CODE_OF_CONDUCT.md
[CONTRIBUTING]: https://github.com/ggoodwin/go-nyse-time/blob/master/CONTRIBUTING.md
[GITHUB-ISSUES]: https://github.com/ggoodwin/go-nyse-time/issues
