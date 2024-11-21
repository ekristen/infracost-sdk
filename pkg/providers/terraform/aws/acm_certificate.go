package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getACMCertificate() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_acm_certificate",
		CoreRFunc: NewACMCertificate,
	}
}
func NewACMCertificate(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ACMCertificate{
		Address:                 d.Address,
		Region:                  d.Get("region").String(),
		CertificateAuthorityARN: d.Get("certificate_authority_arn").String(),
	}
	return r
}
