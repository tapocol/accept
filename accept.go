package accept

import (
	"regexp"
	"net/http"
	"github.com/go-martini/martini"
)

type Format string

const (
	ACCEPT_HEADER = "Accept"
	HTML = Format("html")
	HTML_PATTERN = `(text/html)`
	JSON = Format("json")
	JSON_PATTERN = `(application/json|text/json)`
	XML = Format("xml")
	XML_PATTERN = `(application/xml|text/xml)`
	CSV = Format("csv")
	CSV_PATTERN = `(application/csv|text/csv)`
)

var Formats map[Format]string
var DefaultFormat Format

func init() {
	Formats = make(map[Format]string)
	Formats[HTML] = HTML_PATTERN
	Formats[JSON] = JSON_PATTERN
	DefaultFormat = HTML
}

func Accept(c martini.Context, req *http.Request) {
	accept := req.Header.Get(ACCEPT_HEADER)
	format := DefaultFormat
	if accept != "" {
		for f, p := range Formats {
			found, _ := regexp.MatchString(p, accept)
			if found {
				format = f
				break
			}
		}
	}
	c.Map(&format)
}

