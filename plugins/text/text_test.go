package text

import (
	"testing"

	"github.com/ajinori/d4c/mocks"
)

func TestModule(t *testing.T) {
	c := mocks.Context{}
	if err := Module(c, "message"); err != nil {
		t.Error(err)
	}
}
