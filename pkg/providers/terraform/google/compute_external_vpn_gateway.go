package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getComputeExternalVPNGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_compute_external_vpn_gateway",
		CoreRFunc: NewComputeExternalVPNGateway,
	}
}
func NewComputeExternalVPNGateway(d *schema.ResourceData) schema.CoreResource {
	r := &google.ComputeExternalVPNGateway{Address: d.Address, Region: d.Get("region").String()}
	return r
}
