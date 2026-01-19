package ansible

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/ansible/aws/s3"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws"
	"github.com/deliveroo/trivy/pkg/iac/scanners/ansible/parser"
	"github.com/deliveroo/trivy/pkg/iac/state"
)

func Adapt(tasks parser.ResolvedTasks) state.State {
	return state.State{
		AWS: aws.AWS{
			S3: s3.Adapt(tasks),
			// TODO(simar): Add other AWS services
		},
	}
}
