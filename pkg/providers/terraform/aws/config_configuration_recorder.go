package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getConfigurationRecorderItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_config_configuration_recorder",
		CoreRFunc: NewConfigConfigurationRecorder,
	}
}
func NewConfigConfigurationRecorder(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ConfigConfigurationRecorder{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}
