package dns

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/nifcloud/dns"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) dns.DNS {
	return dns.DNS{
		Records: adaptRecords(modules),
	}
}
