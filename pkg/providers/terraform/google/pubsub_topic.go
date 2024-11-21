package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getPubSubTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_pubsub_topic",
		CoreRFunc: NewPubSubTopic,
	}
}

func NewPubSubTopic(d *schema.ResourceData) schema.CoreResource {
	r := &google.PubSubTopic{
		Address: d.Address,
	}

	return r
}
