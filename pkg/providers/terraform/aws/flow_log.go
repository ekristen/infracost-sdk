package aws

import (
	"github.com/infracost/infracost/pkg/schema"
)

func getFlowLogRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "aws_flow_log",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return schema.BlankCoreResource{
				Name: d.Address,
				Type: d.Type,
			}
		},
	}
}
