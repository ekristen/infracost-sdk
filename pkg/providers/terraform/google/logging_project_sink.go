package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getLoggingProjectSinkRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_project_sink",
		CoreRFunc: NewLoggingProjectSink,
	}
}

func NewLoggingProjectSink(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}
