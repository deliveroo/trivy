package ecs

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/ecs"
	"github.com/deliveroo/trivy/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ECS instance
func Adapt(cfFile parser.FileContext) ecs.ECS {
	return ecs.ECS{
		Clusters:        getClusters(cfFile),
		TaskDefinitions: getTaskDefinitions(cfFile),
	}
}
