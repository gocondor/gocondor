package core_test

import (
	"fmt"
	"testing"

	. "github.com/gincoat/gincoat/core"
)

func TestNew(t *testing.T) {
	app := New()
	typeOfApp := fmt.Sprintf("%T", app)
	if typeOfApp != "*core.App" {
		t.Errorf("expecting app varialbe to be of type *core.App")
	}
}
