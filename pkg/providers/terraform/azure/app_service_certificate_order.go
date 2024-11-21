package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getAppServiceCertificateOrderRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_app_service_certificate_order",
		CoreRFunc: NewAppServiceCertificateOrder,
	}
}
func NewAppServiceCertificateOrder(d *schema.ResourceData) schema.CoreResource {
	r := &azure.AppServiceCertificateOrder{Address: d.Address, ProductType: d.Get("product_type").String()}
	return r
}
