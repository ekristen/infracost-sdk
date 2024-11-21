package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getLightsailInstanceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_lightsail_instance",
		CoreRFunc: NewLightsailInstance,
	}
}

func NewLightsailInstance(d *schema.ResourceData) schema.CoreResource {
	r := &aws.LightsailInstance{
		Address:  d.Address,
		BundleID: d.Get("bundle_id").String(),
		Region:   d.Get("region").String(),
	}
	return r
}
