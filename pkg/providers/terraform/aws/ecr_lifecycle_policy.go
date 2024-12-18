package aws

import (
	"github.com/infracost/infracost/pkg/schema"
)

func getECRLifecyclePolicy() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "aws_ecr_lifecycle_policy",
		ReferenceAttributes: []string{"repository"},
		NoPrice:             true,
		Notes:               []string{"Free resource."},
	}
}
