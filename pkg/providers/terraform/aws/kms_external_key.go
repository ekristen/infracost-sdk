package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getNewKMSExternalKeyRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_kms_external_key",
		CoreRFunc: NewKMSExternalKey,
	}
}

func NewKMSExternalKey(d *schema.ResourceData) schema.CoreResource {
	r := &aws.KMSExternalKey{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}
