# gojsonconf

## Introduction

This package is a simple json configuration decoder into struct.

## Installation

```
go get github.com/petechu/gojsonconf
```

## Usage

### Example 1

#### Configuration file:

```json
// conf.json
{
  "Port": "3000",
  "Host": "example.com"
}
```

#### Code:

```golang
type Conf Struct {
    Port string
    Host string
}

conf := Conf{}
gojsonconf.GetConfig("conf.json", &conf)

fmt.Println(conf.Port) // 3000
fmt.Println(conf.Host) // example.com

```

### Example 2

#### Configuration file:

```json
// conf.json
{
  "port": "3000",
  "host": "example.com"
}
```

#### Code:

```golang
type Conf Struct {
    Port string `json:"port"`
    Host string `json:"host"`
}

conf := Conf{}
gojsonconf.GetConfig("conf.json", &conf)

fmt.Println(conf.Port) // 3000
fmt.Println(conf.Host) // example.com

```
