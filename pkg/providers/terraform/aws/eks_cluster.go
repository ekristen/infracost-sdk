package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getNewEKSClusterItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_eks_cluster",
		CoreRFunc: NewEKSCluster,
	}
}
func NewEKSCluster(d *schema.ResourceData) schema.CoreResource {
	r := &aws.EKSCluster{
		Address: d.Address,
		Region:  d.Get("region").String(),
		Version: d.Get("version").String(),
	}
	return r
}
