package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getNetworkfirewallFirewallRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_networkfirewall_firewall",
		CoreRFunc: newNetworkfirewallFirewall,
	}
}

func newNetworkfirewallFirewall(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()
	r := &aws.NetworkfirewallFirewall{
		Address: d.Address,
		Region:  region,
	}

	return r
}
