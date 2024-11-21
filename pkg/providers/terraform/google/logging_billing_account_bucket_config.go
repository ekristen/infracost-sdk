package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getLoggingBillingAccountBucketConfigRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_billing_account_bucket_config",
		CoreRFunc: NewLoggingBillingAccountBucketConfig,
	}
}

func NewLoggingBillingAccountBucketConfig(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}
