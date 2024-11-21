package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getNATGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "aws_nat_gateway",
		ReferenceAttributes: []string{
			"allocation_id",
			"subnet_id",
		},
		CoreRFunc: NewNATGateway,
	}
}

func NewNATGateway(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()

	a := &aws.NATGateway{
		Address: d.Address,
		Region:  region,
	}

	return a
}
