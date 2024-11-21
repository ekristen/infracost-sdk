package aws

import (
	"github.com/infracost/infracost/pkg/schema"
)

func getOpensearchDomainRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_opensearch_domain",
		CoreRFunc: newSearchDomain,
	}
}
