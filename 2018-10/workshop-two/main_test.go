package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edmontongo/darksky"
)

// mock start OMIT
type forecasterMock struct {
	weather *darksky.Weather
	err     error
}

func (f forecasterMock) Forecast(darksky.Location) (*darksky.Weather, error) {
	return f.weather, f.err
}

// mock end OMIT

// test html start OMIT
func TestHTMLResponse(t *testing.T) {
	forecastMock := forecasterMock{weather: &darksky.Weather{
		Currently: darksky.Currently{Summary: "hot", Temperature: 90},
	}}
	lh := locationHandler{forecastSvc: forecastMock}

	r, _ := http.NewRequest("GET", "/?lat=53.5&lng=-113.3", nil)
	r.Header.Add("Accept", "text/html")
	w := httptest.NewRecorder()
	lh.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("Unexpected status code: %d != 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "hot") {
		t.Errorf("it should be hot")
	}
}

// test html end OMIT
