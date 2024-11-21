package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getDNSPtrRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_ptr_record",
		CoreRFunc: NewDNSPtrRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSPtrRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSPtrRecord{Address: d.Address, Region: d.Region}
	return r
}
