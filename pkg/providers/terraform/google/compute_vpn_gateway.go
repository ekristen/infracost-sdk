package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getComputeVPNGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_compute_vpn_gateway",
		CoreRFunc: NewComputeVPNGateway,
	}
}
func NewComputeVPNGateway(d *schema.ResourceData) schema.CoreResource {
	r := &google.ComputeVPNGateway{Address: d.Address, Region: d.Get("region").String()}
	return r
}
