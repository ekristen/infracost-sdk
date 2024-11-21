package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getDatabricksWorkspaceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_databricks_workspace",
		CoreRFunc: NewDatabricksWorkspace,
	}
}
func NewDatabricksWorkspace(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DatabricksWorkspace{Address: d.Address, Region: d.Region, SKU: d.Get("sku").String()}
	return r
}
