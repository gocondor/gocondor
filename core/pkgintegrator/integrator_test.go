package pkgintegrator_test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/gincoat/gincoat/core/pkgintegrator"
)

func TestNew(t *testing.T) {
	p := New()

	if fmt.Sprintf("%T", p) != "*pkgintegrator.PKGIntegrator" {
		t.Errorf("failed initiate package integrator variable")
	}
}

func TestResolve(t *testing.T) {
	New()
	p := Resolve()

	if fmt.Sprintf("%T", p) != "*pkgintegrator.PKGIntegrator" {
		t.Errorf("failed resolve package integrator variable")
	}
}

func TestIntegrate(t *testing.T) {
	p := New()
	pkgi := func(c *gin.Context) {}
	p.Integrate(pkgi)
	if len(p.GetIntegrations()) != 1 {
		t.Error("failed to integrate a package")
	}
}

func TestGetIntegrations(t *testing.T) {
	p := New()
	p.Integrate(func(c *gin.Context) {})
	p.Integrate(func(c *gin.Context) {})
	if len(p.GetIntegrations()) != 2 {
		t.Error("failed get integrated packages")
	}
}
