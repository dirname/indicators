# Indicators

[![Go Report Card](https://goreportcard.com/badge/github.com/dirname/indicators?style=flat-square)](https://goreportcard.com/report/github.com/dirname/indicators)
[![Build Status](https://img.shields.io/travis/dirname/indicators?style=flat-square)](https://travis-ci.org/dirname/indicators)
[![codecov](https://img.shields.io/codecov/c/gh/dirname/indicators/main?style=flat-square&token=A6U52MYCXN)](https://codecov.io/gh/dirname/indicators)
[![license](https://img.shields.io/github/license/dirname/indicators?style=flat-square)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dirname/indicators?style=flat-square)](https://golang.org)
[![GoDoc](https://img.shields.io/badge/reference-007d9c?style=flat-square&logo=Go&logoColor=F9F9F9&labelColor=5C5C5C&labelWidth=80)](https://pkg.go.dev/github.com/dirname/indicators)
[![Release](https://img.shields.io/github/release/dirname/indicators.svg?style=flat-square)](https://github.com/dirname/indicators/releases)
[![Release Date](https://img.shields.io/github/release-date/dirname/indicators?style=flat-square)](https://github.com/dirname/indicators/releases)
[![Commit](https://img.shields.io/github/last-commit/dirname/indicators?style=flat-square)](https://github.com/dirname/indicators/commits)

The main purpose of the repository is to rewrite some indicator methods
of [trading-indicator](https://gitlab.com/afis/trading-indicator/) into its own real-time trading system to meet the
needs of trading.

# Documents & Wiki

If you need more information about this repository and how to use technical indicators, please read
the [wiki](https://github.com/dirname/indicators/wiki). After reading the wiki, you can quickly use this library and
quote technical indicators.

# Benchmark

- goos: windows
- goarch: amd64
- cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

| Function | Times | Speed
| :----: | :----: | :----:
| BenchmarkMACD_Update-12 | 67301922 | 17.61 ns/op
| BenchmarkRSI_Update-12 | 70786436 | 17.12 ns/op
| BenchmarkMFI_Update-12 | 54278502 | 22.53 ns/op
| BenchmarkSMA_Update-12 | 58746352 | 17.24 ns/op
| BenchmarkVariance_Update-12 | 44447077 | 23.24 ns/op
| BenchmarkVariance_Sum-12 | 44626918 | 24.86 ns/op
| BenchmarkDEMA_Update-12 | 61096374 | 19.85 ns/op
| BenchmarkEMA_Update-12 | 164322492 | 7.148 ns/op
| BenchmarkTEMA_Update-12 | 44629740 | 25.66 ns/op
| BenchmarkWMA_Update-12 | 50345283 | 22.75 ns/op
| BenchmarkVariance_Update-12 | 51031256 | 21.86 ns/op
| BenchmarkStdDev_Sum-12 | 51484468 | 22.48 ns/op
| BenchmarkMAMA_Update-12 | 14454295 | 82.35 ns/op
| BenchmarkTRIMA_Update-12 | 52316760 | 21.39 ns/op
| BenchmarkKAMA_Update-12 | 44349830 | 24.15 ns/op
| BenchmarkT3MA_Update-12 | 69494372 | 16.94 ns/op
| BenchmarkBands_Update-12 | 27291334 | 39.07 ns/op
| BenchmarkCross_CrossOver-12 | 651629235 | 1.799 ns/op
| BenchmarkCross_CrossUnder-12 | 643732273 | 1.815 ns/op

# Disclaimer

The purpose of this repository is to learn technical indicators, **not** in the *production environment*, so some
technical indicators may be wrong under certain conditions. Please use it with caution, and expand the R&D according to
your own situation.

**In addition**, the performance and stability of this repository and the correctness of the technical indicators need
to be tested by yourself, and the losses caused by the performance or stability and the correctness of the indicators
should be **borne by you**.
