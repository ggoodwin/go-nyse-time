<div align="center">
	<h1><img alt="Stocks logo" src="https://github.com/octolibs/nyse-time/blob/main/timetable.png" height="300" /><br />
		New York Stock Exchange (NYSE) Time Go Library
	</h1>

[![Go Reference](https://pkg.go.dev/badge/octolibs/nyse-time.svg)](https://pkg.go.dev/github.com/octolibs/nyse-time) [![Go Version](https://img.shields.io/github/go-mod/go-version/octolibs/nyse-time)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/nyse-time)](https://goreportcard.com/report/github.com/octolibs/nyse-time) [![CodeFactor](https://www.codefactor.io/repository/github/octolibs/nyse-time/badge)](https://www.codefactor.io/repository/github/octolibs/nyse-time) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/octolibs/nyse-time/.github/workflows/go.yml)](https://github.com/octolibs/nyse-time/blob/main/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/octolibs/nyse-time) [![Last Commit](https://img.shields.io/github/last-commit/octolibs/nyse-time)](https://github.com/octolibs/nyse-time/commits/main) [![License](https://img.shields.io/github/license/octolibs/nyse-time)](https://github.com/octolibs/nyse-time/blob/main/LICENSE)

</div>
<hr/>

## 🌟 How it works

The library uses the current time to determine if the market is open, closed, or early closed. It also determines if the current day is a holiday.

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
go get github.com/octolibs/nyse-time
```

### Importing

Add the import to your `.go` file

```go
import "github.com/octolibs/nyse-time"
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

### Is Holiday?

```go
// Parameters: time.Time
// Returns: bool
isHoliday := nyse_time.IsHoliday(time.Time)
```

### Is Early Close?

```go
// Parameters: time.Time
// Returns: bool
isEarlyClose := nyse_time.IsEarlyClose(time.Time)
```

### Is Weekend?

```go
// Parameters: time.Time
// Returns: bool
isWeekend := nyse_time.IsWeekend(time.Time)
```

## 💻 Dependencies

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.

## ⚖️ License

This project is under the MIT License. See the [LICENSE](https://github.com/octolibs/nyse-time/blob/main/LICENSE) file for the full license text.

## 📜 Changes

Check out our [CHANGELOG](https://github.com/octolibs/nyse-time/blob/main/CHANGELOG.md)
