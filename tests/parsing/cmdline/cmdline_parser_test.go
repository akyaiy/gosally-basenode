package cmdline_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/akyaiy/gosally-basenode/internal/config"
	"github.com/akyaiy/gosally-basenode/internal/parser"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestParseArgs(t *testing.T) {
	args := []string{"prog", "--listen-address=0.0.0.0"}

	var p = &parser.Parser{Config: config.Config{}}
	err := p.ParseArgs(args)
	if err != nil {
		t.Errorf("ParseArgs failed: %v", err)
	}
	want := config.TypeConfig{
		"exec-name":      config.ConfValue{Value: "prog", Source: config.SrcCMD},
		"listen-address": config.ConfValue{Value: "0.0.0.0", Source: config.SrcCMD},
	}
	if !reflect.DeepEqual(want, p.Config.Vars) {
		t.Fatalf("Expected %+v, got %+v", want, p.Config.Vars)
	}
}
