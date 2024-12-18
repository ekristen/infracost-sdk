package azure_test

import (
	"testing"

	"github.com/infracost/infracost/pkg/providers/terraform/tftest"
)

func TestMonitorDataCollectionRule(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "monitor_data_collection_rule_test")
}
