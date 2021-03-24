package core_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gincoat/gincoat/config"
	"github.com/gincoat/gincoat/core"
	. "github.com/gincoat/gincoat/core"
	"github.com/gincoat/gincoat/core/database"
	"github.com/gincoat/gincoat/core/middlewaresengine"
	"github.com/gincoat/gincoat/core/pkgintegrator"
	"github.com/gincoat/gincoat/core/routing"
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

func TestRegisterRoutes(t *testing.T) {
	routes := []routing.Route{
		{
			Method: "get",
			Path:   "/:name",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
					val, _ := c.Params.Get("name")
					c.JSON(http.StatusOK, gin.H{
						"name": val,
					})
				},
			},
		},
	}
	g := gin.New()
	app := New()
	app.RegisterRoutes(routes, g)
	s := httptest.NewServer(g)
	defer s.Close()

	res, _ := http.Get(fmt.Sprintf("%s/jack", s.URL))
	body, _ := ioutil.ReadAll(res.Body)
	type ResultStruct struct {
		Name string `json:"Name"`
	}
	var result ResultStruct
	json.Unmarshal(body, &result)
	if result.Name != "jack" {
		t.Errorf("failed assert execution of registerd route")
	}
}

func TestSetEnabledFeatures(t *testing.T) {
	app := New()
	app.SetEnabledFeatures(config.Features)

	if app.Features.Database != false || app.Features.Cache != false || app.Features.GRPC != false {
		t.Errorf("failed setting features")
	}
}

func TestBootstrap(t *testing.T) {
	var Features *core.Features = &core.Features{
		Database: true,
		Cache:    false,
		GRPC:     false,
	}
	os.Setenv("DB_DRIVER", "sqlite") // set database driver to sqlite
	app := New()
	env, _ := godotenv.Read("./testingdata/.env")
	app.SetEnv(env)
	app.SetEnabledFeatures(Features)
	app.Bootstrap()

	i := pkgintegrator.Resolve()
	if i == nil || fmt.Sprintf("%T", i) != "*pkgintegrator.PKGIntegrator" {
		t.Errorf("failed asserting the initiation of PKGIntegrator")
	}

	m := middlewaresengine.Resolve()
	if m == nil || fmt.Sprintf("%T", m) != "*middlewaresengine.MiddlewaresEngine" {
		t.Errorf("failed asserting the initiation of MiddlewaresEngine")
	}

	r := routing.Resolve()
	if r == nil || fmt.Sprintf("%T", r) != "*routing.Router" {
		t.Errorf("failed asserting the initiation of Router")
	}

	d := database.Resolve()
	if d == nil || fmt.Sprintf("%T", d) != "*gorm.DB" {
		t.Errorf("failed asserting the initiation of Database")
	}
}

func TestUseMiddleWares(t *testing.T) {
	middlewares := []gin.HandlerFunc{
		func(c *gin.Context) {
			c.Set("VAR1", "VAL1")
		},
		func(c *gin.Context) {
			c.Set("VAR2", "VAL2")
		},
	}

	g := gin.New()
	app := New()
	g = app.UseMiddlewares(middlewares, g)

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"VAR1": c.MustGet("VAR1"),
			"VAR2": c.MustGet("VAR2"),
		})
	})
	s := httptest.NewServer(g)
	defer s.Close()
	res, _ := s.Client().Get(s.URL)
	body, _ := ioutil.ReadAll(res.Body)

	type ResponseStruct struct {
		VAR1 string `json:"VAR1"`
		VAR2 string `json:"VAR2"`
	}

	var response ResponseStruct
	json.Unmarshal(body, &response)

	if response.VAR1 != "VAL1" || response.VAR2 != "VAL2" {
		t.Errorf("failed asserting middlewares registering")
	}
}

func TestGetHTTPSHost(t *testing.T) {
	app := New()
	host := app.GetHTTPSHost()
	if host != "localhost" {
		t.Errorf("failed getting https host")
	}

	os.Setenv("APP_HTTPS_HOST", "testserver.com")
	host = app.GetHTTPSHost()
	if host != "testserver.com" {
		t.Errorf("failed getting https host")
	}

}

//GetHTTPSHost
//getHTTPHost
//Run
