Accept (Martini Middleware)
===========================

Using regexp patterns in Go, this project will match the "Accept" HTTP Header to a Format string.
The Format string will be mapped during the Martini Middleware.

Usage
-----
Every Request:

```go
m.Use(accept.Accept)
m.Get("/", func(f *Format) string {
  switch *f {
  case accept.JSON:
    return `{"Hello": "JSON"}`
  default:
    return "Hello HTML"
  }
})
```

Route-specific:

```go
m.Get("/", accept.Accept, func(f *Format) string {
  switch *f {
  case accept.JSON:
    return `{"Hello": "JSON"}`
  default:
    return "Hello HTML"
  }
})
```

Setup
-----

Download:

```sh
$ go get github.com/craigjackson/accept
```

Change the default format (inititalizes to accept.HTML):

```go
accept.DefaultFormat = accept.JSON
```

Add addition formats to attempt to match (initializes to just use HTML and JSON):

```go
accept.Formats[XML] = accept.XML_PATTERN
```

