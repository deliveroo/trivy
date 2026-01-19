package terraform

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/aws"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/azure"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/cloudstack"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/digitalocean"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/github"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/google"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/kubernetes"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/nifcloud"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/openstack"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/oracle"
	"github.com/deliveroo/trivy/pkg/iac/state"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}
