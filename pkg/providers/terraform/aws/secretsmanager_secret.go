package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getSecretsManagerSecret() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_secretsmanager_secret",
		CoreRFunc: NewSecretsManagerSecret,
	}
}

func NewSecretsManagerSecret(d *schema.ResourceData) schema.CoreResource {
	r := &aws.SecretsManagerSecret{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}
