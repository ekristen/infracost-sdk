package azure_test

import (
	"testing"

	"github.com/infracost/infracost/pkg/providers/terraform/tftest"
)

func TestImage(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTestsWithOpts(t, "image_test", &tftest.GoldenFileOptions{IgnoreCLI: true})
}
