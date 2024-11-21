package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getCloudFormationStackRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_cloudformation_stack",
		CoreRFunc: NewCloudFormationStackSet,
	}
}
func NewCloudFormationStack(d *schema.ResourceData) schema.CoreResource {
	r := &aws.CloudFormationStack{
		Address:      d.Address,
		Region:       d.Get("region").String(),
		TemplateBody: d.Get("template_body").String(),
	}
	return r
}
