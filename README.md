# iconvert

[![Go Report Card](https://goreportcard.com/badge/github.com/gophreak/iconvert)](https://goreportcard.com/report/github.com/gophreak/iconvert)

Converts from an interface into one of the supported native go types:

* String
* Float64
* Int64

## Getting started

### go get

Run this command: `go get github.com/gophreak/iconvert`

### dep

Run this command: `dep ensure -add github.com/gophreak/iconvert` to install the package.

### glide

Download the package using glide by adding the following into your glide.yaml file

```yml
import:
- package: github.com/gophreak/iconvert
  repo: git@github.com:gophreak/iconvert.git
  vcs: git
```

Then run `glide up -v` to install the package.

## Using the package

The following functions are available for use:

* `ToString(value interface{})` which takes an interface and returns a string representation of it
* `ToFloat(value interface{})` which takes an interface and returns a float64 representation of it
* `ToInt(value interface{})` which takes an interface and returns an int64 representation of it

In both cases if the conversion is not possible, then an error is returned. In the event of calling ToFloat with a string
which cannot be a float, the word "Hello" for example, then the native error from the go library will be returned.

## Examples

To convert the integer value of `6` to a `string`, run the following:

```go
v, e := iconvert.ToString(int(6))
```

`v` should equal `"6"`, and `e` should be `nil`.

To convert the value of the string `"6"` to a `float64`, run the following:

```go
v, e := iconvert.ToFloat("6")
```

`v` should equal `float64(6)`, and `e` should be `nil`

To convert the value of the float `6.35` to an `integer`, run the following:

```go
v, e := iconvert.ToInt("6.35")
```

The decimal places will be removed, and `v` should equal `int64(6)`, and `e` should be `nil`
