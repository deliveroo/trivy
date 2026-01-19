package nifcloud

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/computing"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/dns"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/nas"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/network"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/rdb"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud/sslcertificate"
	"github.com/deliveroo/trivy/pkg/iac/providers/nifcloud"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) nifcloud.Nifcloud {
	return nifcloud.Nifcloud{
		Computing:      computing.Adapt(modules),
		DNS:            dns.Adapt(modules),
		NAS:            nas.Adapt(modules),
		Network:        network.Adapt(modules),
		RDB:            rdb.Adapt(modules),
		SSLCertificate: sslcertificate.Adapt(modules),
	}
}
