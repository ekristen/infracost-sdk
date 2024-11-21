package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getELBRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_elb",
		CoreRFunc: NewELB,
	}
}
func NewELB(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ELB{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}
