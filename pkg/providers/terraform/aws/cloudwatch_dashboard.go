package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getCloudwatchDashboardRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_cloudwatch_dashboard",
		CoreRFunc: NewCloudwatchDashboard,
	}
}
func NewCloudwatchDashboard(d *schema.ResourceData) schema.CoreResource {
	r := &aws.CloudwatchDashboard{
		Address: d.Address,
	}
	return r
}
