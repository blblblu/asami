# asami

[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

pixel sorter using simple brute sorting

## installation

`go get -u github.com/blblblu/asami`

## usage

```
> asami --help
```

```
simple image corrupter

Usage:
  asami [command]

Available Commands:
  sort        simple brute pixel sorting

Use "asami [command] --help" for more information about a command.
```

```
asami sort --help
```

```
simple brute pixel sorting

Usage:
  asami sort [flags]

Flags:
  -i, --input string    the input file path to use
      --inverted        inverts the sorting direction
      --max int         the maximum chunk size to use (default 64)
      --min int         the minimum chunk size to use (default 32)
  -o, --output string   the output file path to use, must be a png file
```
