# web
Go web utilities

## Webserve

Start a webserver concurrently

### Install

`go get github.com/whitecypher/web`

### Usage

```go
r := http.NewServerMux()

// ... add handlers to servermux

s := webserve.New(":8080", r)
s.Start()

// ... any other code

// Use select when running multiple servers and exit when one of them fails
select {
	case err := s.Wait():
		log.Fatal(err)
}
```