package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getAPIGatewayStageRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_api_gateway_stage",
		CoreRFunc: NewAPIGatewayStage,
	}
}
func NewAPIGatewayStage(d *schema.ResourceData) schema.CoreResource {
	r := &aws.APIGatewayStage{
		Address:          d.Address,
		Region:           d.Get("region").String(),
		CacheClusterSize: d.Get("cache_cluster_size").Float(),
		CacheEnabled:     d.GetBoolOrDefault("cache_cluster_enabled", false),
	}
	return r
}
