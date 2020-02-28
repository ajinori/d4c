package config

import (
	"net/http"
)

var (
	Default = Config{
		Routes: []Route{
			Route{
				Path:   "/",
				Method: http.MethodGet,
				Text: Text{
					Message: "d4c",
				},
			},
		},
	}
)

type Config struct {
	Routes []Route
}

type Route struct {
	Path   string `json:"path" yaml:"path"`
	Method string `json:"method" yaml:"method"`
	Text   Text   `json:"text" yaml:"text"`
}

type Text struct {
	Message string `json:"message" yaml:"message"`
}
