package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getDNSManagedZoneRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_dns_managed_zone",
		CoreRFunc: NewDNSManagedZone,
	}
}

func NewDNSManagedZone(d *schema.ResourceData) schema.CoreResource {
	r := &google.DNSManagedZone{
		Address: d.Address,
	}

	return r
}
