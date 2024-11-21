package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getEventgridTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_eventgrid_topic",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return &azure.EventGridTopic{
				Address: d.Address,
				Region:  d.Region,
			}
		},
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
