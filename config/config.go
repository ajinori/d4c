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
				Plugin: Plugin{
					Name:    "text",
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
	Plugin Plugin `json:"plugin" yaml:"plugin"`
}

type Plugin struct {
	Name    string `json:"name" yaml:"name"`
	Message string `json:"message" yaml:"message"`
}
