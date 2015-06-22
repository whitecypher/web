# web
Go web utilities

## Webserve

Start a webserver concurrently

### Install

`go get github.com/whitecypher/web`

### Usage

```go
r := http.NewServerMux()
s := webserve.New(":8080", r)
s.Start()

// ...

select {
	case err := s.Wait():
		log.Fatal(err)
}
```