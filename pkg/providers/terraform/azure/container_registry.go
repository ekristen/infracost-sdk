package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getContainerRegistryRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_container_registry",
		CoreRFunc: NewContainerRegistry,
	}
}
func NewContainerRegistry(d *schema.ResourceData) schema.CoreResource {
	r := &azure.ContainerRegistry{
		Address:                 d.Address,
		Region:                  d.Region,
		GeoReplicationLocations: len(d.Get("georeplications").Array()),
		SKU:                     d.Get("sku").String(),
	}
	return r
}
