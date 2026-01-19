package dynamodb

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/dynamodb"
	"github.com/deliveroo/trivy/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a dynamodb instance
func Adapt(cfFile parser.FileContext) dynamodb.DynamoDB {
	return dynamodb.DynamoDB{
		DAXClusters: getClusters(cfFile),
		Tables:      getTables(cfFile),
	}
}
