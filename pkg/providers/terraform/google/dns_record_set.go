package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getDNSRecordSetRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_dns_record_set",
		CoreRFunc: NewDNSRecordSet,
	}
}

func NewDNSRecordSet(d *schema.ResourceData) schema.CoreResource {
	r := &google.DNSRecordSet{
		Address: d.Address,
	}

	return r
}
