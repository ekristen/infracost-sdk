package aws_test

import (
	"testing"

	"github.com/infracost/infracost/pkg/providers/terraform/tftest"
)

func TestBackupVault(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "backup_vault_test")
}