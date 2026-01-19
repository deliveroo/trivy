package digitalocean

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/digitalocean/compute"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/digitalocean/spaces"
	"github.com/deliveroo/trivy/pkg/iac/providers/digitalocean"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) digitalocean.DigitalOcean {
	return digitalocean.DigitalOcean{
		Compute: compute.Adapt(modules),
		Spaces:  spaces.Adapt(modules),
	}
}
