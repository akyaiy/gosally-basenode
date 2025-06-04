package database

import (
	"reflect"
	"testing"

	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func TestConnectionBuilder_Build(t *testing.T) {
	opts := NewConnection().
		WithConnectionString("file::memory:?cache=shared").
		WithTimeout(30).
		WithConnectionID("test-connection-id").
		EndBuild()

	expectedOpts := &DatabaseConnectionOpt{
		ConnectionString: "file::memory:?cache=shared",
		Timeout:          30,
		ConnectionID:     "test-connection-id",
	}

	if !reflect.DeepEqual(opts, expectedOpts) {
		t.Errorf("Expected opts: %+v, got: %+v", expectedOpts, opts)
	}
}

func TestDriverBuilder_Build(t *testing.T) {
	log := logger.NewMockLogger()
	driver := NewDriver().
		WithLogger(log).
		WithDriverType(DriverTypeSQLite).
		EndBuild()
	
	if driver.Log != &log {
		t.Errorf("Expected logger: %v, got: %v", &log, driver.Log)
	}
	if driver.driver != DriverTypeSQLite {
		t.Errorf("Expected driver type: %d, got: %d", DriverTypeSQLite, driver.driver)
	}
}