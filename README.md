<div align="center">
	<h1><img alt="Stocks logo" src="https://github.com/octolibs/calendar/blob/main/timetable.png" height="300" /><br />
		Go New York Stock Exchange (NYSE) Schedule Library
	</h1>

[![Go Reference](https://pkg.go.dev/badge/octolibs/calendar.svg)](https://pkg.go.dev/github.com/octolibs/calendar) [![Go Version](https://img.shields.io/github/go-mod/go-version/octolibs/calendar)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/octolibs/calendar)](https://goreportcard.com/report/github.com/octolibs/calendar) [![CodeFactor](https://www.codefactor.io/repository/github/octolibs/calendar/badge)](https://www.codefactor.io/repository/github/octolibs/calendar) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/octolibs/calendar/.github/workflows/go.yml)](https://github.com/octolibs/calendar/blob/main/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/octolibs/calendar) [![Last Commit](https://img.shields.io/github/last-commit/octolibs/calendar)](https://github.com/octolibs/calendar/commits/main) [![License](https://img.shields.io/github/license/octolibs/calendar)](https://github.com/octolibs/calendar/blob/main/LICENSE)

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
go get github.com/octolibs/nyse-schedule
```

### Importing

Add the import to your `.go` file

```go
import "github.com/octolibs/nyse-schedule"
```

## 💰 Usage

### Check if the Market is Open

```go
//Returns True or False
open := nyse.IsMarketOpen()
```

### Get full stock details

```go
//Returns `Result` struct
stock := stocks.GetFullDetails("AAPL")
```

### Is Market Open?

```go
//Returns `bool`
isOpen := stocks.IsMarketOpen()
```

### Is Holiday?

```go
//Returns `bool`
isHoliday := stocks.IsHoliday()
```

### Is Early Close?

```go
//Returns `bool`
isEarlyClose := stocks.IsEarlyClose()
```

## 💻 Dependencies

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.

## ⚖️ License

This project is under the MIT License. See the [LICENSE](https://github.com/octolibs/calendar/blob/main/LICENSE) file for the full license text.

## 📜 Changes

Check out our [CHANGELOG](https://github.com/octolibs/calendar/blob/main/CHANGELOG.md)
