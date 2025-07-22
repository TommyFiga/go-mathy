# go-mathy
> A simple CLI calculator written in Go

`go-mathy` is a lightweight command-line tool that performs basic arithmetic operations.
It was created as a learning project while exploring Go and CLI development.

---

## Features
- Basic math operations: `+`, `-`, `*`, `/`
- Supports `x` or `X` for multiplication
- Friendly error handling
- Cross-platform binary (build it anywhere)

---

## Installation
### Option 1: `go install` (requires Go 1.18+)

``` bash
go install github.com/TommyFiga/go-mathy@latest
```

> **Note:** The binary is installed to `$GOPATH/bin` (usually `~/go/bin`).

### Option 2: Clone & Build

``` bash
git clone https://github.com/TommyFiga/go-mathy.git
cd go-mathy
go build -o go-mathy
```

Then (_optional_), move the binary to a system-wide location:

``` bash
sudo mv go-mathy /usr/local/bin/
```

---

## Usage

``` bash
go-mathy <number1> <operator> <number2>
```

### Example:
``` bash
go-mathy 5 + 3
```

``` bash
go-mathy 10 / 2
```

---

## Multiplication Note
When using `*`, you must wrap it in quotes to prevent shell expansion. 

``` bash
go-mathy 4 "*" 3
```

Alternatively, `go-mathy` supports the usage of `x` and `X` for the multiplication operation.

``` bash
go-mathy 4 x 3
go-mathy 4 X 3
```