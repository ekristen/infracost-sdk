package aws

import (
	"github.com/infracost/infracost/pkg/schema"
)

func getCloudwatchEventTargetRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "aws_cloudwatch_event_target",
		ReferenceAttributes: []string{"ecs_target.0.task_definition_arn"},
		NoPrice:             true,
		Notes:               []string{"Free resource."},
	}
}
