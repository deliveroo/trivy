package cloudstack

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/cloudstack/compute"
	"github.com/deliveroo/trivy/pkg/iac/providers/cloudstack"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) cloudstack.CloudStack {
	return cloudstack.CloudStack{
		Compute: compute.Adapt(modules),
	}
}
