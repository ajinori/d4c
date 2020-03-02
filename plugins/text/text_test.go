package text

import (
	"testing"

	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/mocks"
)

func TestModule(t *testing.T) {
	c := mocks.Context{}
	cfg := config.Route{
		Plugin: config.Plugin{
			Message: "message",
		},
	}

	if err := Run(c, cfg); err != nil {
		t.Error(err)
	}
}
