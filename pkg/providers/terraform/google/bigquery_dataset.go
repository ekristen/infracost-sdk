package google

import (
	"github.com/infracost/infracost/pkg/resources/google"
	"github.com/infracost/infracost/pkg/schema"
)

func getBigQueryDatasetRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_bigquery_dataset",
		CoreRFunc: NewBigQueryDataset,
	}
}

func NewBigQueryDataset(d *schema.ResourceData) schema.CoreResource {
	r := &google.BigQueryDataset{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return r
}
