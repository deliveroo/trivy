package digitalocean

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/digitalocean/compute"
	"github.com/deliveroo/trivy/pkg/iac/providers/digitalocean/spaces"
)

type DigitalOcean struct {
	Compute compute.Compute
	Spaces  spaces.Spaces
}
