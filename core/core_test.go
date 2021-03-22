package core_test

import (
	"fmt"
	"os"
	"testing"

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
