package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getPrivateDNSCNameRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_private_dns_cname_record",
		CoreRFunc: NewPrivateDNSCNameRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewPrivateDNSCNameRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.PrivateDNSCNameRecord{Address: d.Address, Region: d.Region}
	return r
}
