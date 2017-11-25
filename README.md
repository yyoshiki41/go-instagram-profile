# go-instagram-profile

[![godoc](https://godoc.org/github.com/yyoshiki41/go-instagram-profile?status.svg)](https://godoc.org/github.com/yyoshiki41/go-instagram-profile)
[![build](https://travis-ci.org/yyoshiki41/go-instagram-profile.svg?branch=master)](https://travis-ci.org/yyoshiki41/go-instagram-profile)

Access instagram's public data __without access_token__

## About

A user's profile on instagram can be accessed on the below URL.

```bash
https://www.instagram.com/<username>/?__a=1
```

### Curl Example

```bash
$ curl "https://www.instagram.com/natgeo/?__a=1" | jq .user.biography
"Experience the world through the eyes of National Geographic photographers."
```

## Usage

### API

- `GetProfile()`

That's all !

```go
package main

import (
	"log"

	"github.com/yyoshiki41/go-instagram-profile"
)

func main() {
	u, err := instagram.GetProfile("natgeo")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", u)
}
```

## License 

The MIT License

## Author

Yoshiki Nakagawa
