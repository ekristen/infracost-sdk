package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getNetworkWatcherRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_network_watcher",
		CoreRFunc: newNetworkWatcher,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newNetworkWatcher(d *schema.ResourceData) schema.CoreResource {
	region := d.Region
	return &azure.NetworkWatcher{
		Address: d.Address,
		Region:  region,
	}
}
