package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getExpressRouteConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_express_route_connection",
		RFunc: newExpressRouteConnection,
	}
}

func newExpressRouteConnection(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	e := &azure.ExpressRouteConnection{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return e.BuildResource()
}
