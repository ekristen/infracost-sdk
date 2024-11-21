package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getRoute53ZoneRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_route53_zone",
		CoreRFunc: NewRoute53Zone,
	}
}

func NewRoute53Zone(d *schema.ResourceData) schema.CoreResource {
	r := &aws.Route53Zone{
		Address: d.Address,
	}
	return r
}
