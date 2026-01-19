package cloudformation

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/cloudformation/aws"
	"github.com/deliveroo/trivy/pkg/iac/scanners/cloudformation/parser"
	"github.com/deliveroo/trivy/pkg/iac/state"
)

// Adapt adapts the Cloudformation instance
func Adapt(cfFile parser.FileContext) *state.State {
	return &state.State{
		AWS: aws.Adapt(cfFile),
	}
}
