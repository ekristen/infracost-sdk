package azure_test

import (
	"testing"

	"github.com/infracost/infracost/pkg/providers/terraform/tftest"
)

func TestLogicAppStandard(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "logic_app_standard_test")
}
