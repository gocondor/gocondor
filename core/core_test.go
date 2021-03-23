package core_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/gincoat/gincoat/core"
	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	app := New()
	typeOfApp := fmt.Sprintf("%T", app)
	if typeOfApp != "*core.App" {
		t.Errorf("expecting app varialbe to be of type *core.App")
	}
}

func TestSetEnv(t *testing.T) {
	env, err := godotenv.Read("./testingdata/.env")
	if err != nil {
		t.Errorf("failed reading .env file")
	}
	app := New()
	app.SetEnv(env)

	if os.Getenv("KEY_ONE") != "VAL_ONE" || os.Getenv("KEY_TWO") != "VAL_TWO" {
		t.Errorf("failed to set env vars")
	}
}

func TestSetAppMode(t *testing.T) {
	app := New()
	app.SetAppMode("release")

	if gin.Mode() != gin.ReleaseMode {
		t.Errorf("failed to set app mode")
	}
}

func TestIntegratePackages(t *testing.T) {
	g := gin.New()
	app := New()
	hanldrFuncs := []gin.HandlerFunc{
		func(c *gin.Context) {
			c.Set("TEST_KEY1", "TEST_VAL1")
		},
		func(c *gin.Context) {
			c.Set("TEST_KEY2", "TEST_VAL2")
		},
	}
	g = app.IntegratePackages(g, hanldrFuncs)
	g.GET("/", func(c *gin.Context) {
		if c.MustGet("TEST_KEY1") != "TEST_VAL1" || c.MustGet("TEST_KEY2") != "TEST_VAL2" {
			t.Errorf("failed to integrate packages")
		}
	})

	server := httptest.NewServer(g)
	defer server.Close()
	_, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}
}
