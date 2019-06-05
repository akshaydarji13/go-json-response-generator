# go-json-response-generator
Generates json string for sending it as a response to a network request.

## Installation

```
go get github.com/akshaydarji13/go-json-response-generator
```

## Usage

```
import (
  generator "github.com/akshaydarji13/go-json-response-generator"
)
generator.ResponseGenerator(true, "Signup successful", data)
```

## Response String Format
```
{"status":true,"message":"response message","data":"[{"data":"response data"}]"}
```

## `ResponseGenerator` function parameters

**status**: `bool`
**message**: `string`
**data**: `interface{}`(optional)


