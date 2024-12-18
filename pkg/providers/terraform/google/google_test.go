package google_test

import (
	"os"
	"testing"

	"github.com/infracost/infracost/pkg/providers/terraform/tftest"
)

func TestMain(m *testing.M) {
	tftest.EnsurePluginsInstalled()
	code := m.Run()
	os.Exit(code)
}
