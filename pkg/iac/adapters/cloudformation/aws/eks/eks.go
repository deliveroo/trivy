package eks

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/eks"
	"github.com/deliveroo/trivy/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an EKS instance
func Adapt(cfFile parser.FileContext) eks.EKS {
	return eks.EKS{
		Clusters: getClusters(cfFile),
	}
}
