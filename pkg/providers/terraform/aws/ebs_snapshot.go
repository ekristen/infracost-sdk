package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getEBSSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "aws_ebs_snapshot",
		CoreRFunc:           NewEBSSnapshot,
		ReferenceAttributes: []string{"volume_id"},
	}
}
func NewEBSSnapshot(d *schema.ResourceData) schema.CoreResource {
	r := &aws.EBSSnapshot{Address: d.Address, Region: d.Get("region").String()}
	volumeRefs := d.References("volume_id")
	if len(volumeRefs) > 0 {
		if volumeRefs[0].Get("size").Exists() {
			r.SizeGB = floatPtr(volumeRefs[0].Get("size").Float())
		}
	}
	return r
}
