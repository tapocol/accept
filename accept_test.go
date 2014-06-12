package accept

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/go-martini/martini"
)

func Test_Accept_Default(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != HTML {
			t.Error("Should get HTML format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(rec, req)
}

func Test_Accept_OtherDefault(t *testing.T) {
	DefaultFormat = JSON
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != JSON {
			t.Error("Should get JSON format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(rec, req)
}

func Test_Accept_Text_HTML(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != HTML {
			t.Error("Should get HTML format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "text/html")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Application_JSON(t *testing.T) {
	DefaultFormat = HTML
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != JSON {
			t.Error("Should get JSON format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "application/json")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Text_JSON(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != JSON {
			t.Error("Should get JSON format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "text/json")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Application_XML(t *testing.T) {
	Formats[XML] = XML_PATTERN
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != XML {
			t.Error("Should get XML format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "application/xml")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Text_XML(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != XML {
			t.Error("Should get XML format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "text/xml")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Application_CSV(t *testing.T) {
	Formats[CSV] = CSV_PATTERN
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != CSV {
			t.Error("Should get CSV format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "application/csv")
	m.ServeHTTP(rec, req)
}

func Test_Accept_Text_CSV(t *testing.T) {
	m := martini.Classic()
	m.Get("/", Accept, func(f *Format) {
		if *f != CSV {
			t.Error("Should get CSV format")
		}
	})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(ACCEPT_HEADER, "text/csv")
	m.ServeHTTP(rec, req)
}
