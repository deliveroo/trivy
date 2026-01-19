package sslcertificate

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/nifcloud/sslcertificate"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) sslcertificate.SSLCertificate {
	return sslcertificate.SSLCertificate{
		ServerCertificates: adaptServerCertificates(modules),
	}
}
