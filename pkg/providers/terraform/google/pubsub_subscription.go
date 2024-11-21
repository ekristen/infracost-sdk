package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getPubSubSubscriptionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_pubsub_subscription",
		CoreRFunc: NewPubSubSubscription,
	}
}

func NewPubSubSubscription(d *schema.ResourceData) schema.CoreResource {
	r := &google.PubSubSubscription{
		Address: d.Address,
	}

	return r
}
