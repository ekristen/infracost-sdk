package azure

import (
	"github.com/infracost/infracost/pkg/resources/azure"
	"github.com/infracost/infracost/pkg/schema"
)

func getMachineLearningComputeInstanceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_machine_learning_compute_instance",
		CoreRFunc: newMachineLearningComputeInstance,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newMachineLearningComputeInstance(d *schema.ResourceData) schema.CoreResource {
	region := d.Region
	return &azure.MachineLearningComputeInstance{
		Address:      d.Address,
		Region:       region,
		InstanceType: d.Get("virtual_machine_size").String(),
	}
}
