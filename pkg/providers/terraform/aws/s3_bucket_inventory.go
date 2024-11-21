package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getS3BucketInventoryRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_s3_bucket_inventory",
		CoreRFunc: NewS3BucketInventory,
	}
}

func NewS3BucketInventory(d *schema.ResourceData) schema.CoreResource {
	r := &aws.S3BucketInventory{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}
