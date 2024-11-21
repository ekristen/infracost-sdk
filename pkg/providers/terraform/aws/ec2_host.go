package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getEC2HostRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_ec2_host",
		CoreRFunc: newEC2Host,
	}
}

func newEC2Host(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()
	r := &aws.EC2Host{
		Address:        d.Address,
		Region:         region,
		InstanceType:   d.Get("instance_type").String(),
		InstanceFamily: d.Get("instance_family").String(),
	}
	return r
}
