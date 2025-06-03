package builder_test

import (
	"reflect"
	"testing"

	"github.com/akyaiy/gosally-basenode/internal/database"
)

func TestConnectionBuilder_Build(t *testing.T) {
	opts := database.NewConnection().
		WithConnectionString("file::memory:?cache=shared").
		WithTimeout(30).
		WithConnectionID("test-connection-id").
		EndBuild()

	expectedOpts := &database.DatabaseConnectionOpt{
		ConnectionString: "file::memory:?cache=shared",
		Timeout:          30,
		ConnectionID:     "test-connection-id",
	}

	if !reflect.DeepEqual(opts, expectedOpts) {
		t.Errorf("Expected opts: %+v, got: %+v", expectedOpts, opts)
	}
}
