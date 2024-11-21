package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getDMSRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_dms_replication_instance",
		CoreRFunc: NewDMSReplicationInstance,
	}
}

func NewDMSReplicationInstance(d *schema.ResourceData) schema.CoreResource {
	r := &aws.DMSReplicationInstance{
		Address:                  d.Address,
		MultiAZ:                  d.Get("multi_az").Bool(),
		AllocatedStorageGB:       d.Get("allocated_storage").Int(),
		ReplicationInstanceClass: d.Get("replication_instance_class").String(),
		Region:                   d.Get("region").String(),
	}
	return r
}
