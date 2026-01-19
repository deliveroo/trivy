package github

import (
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/github/branch_protections"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/github/repositories"
	"github.com/deliveroo/trivy/pkg/iac/adapters/terraform/github/secrets"
	"github.com/deliveroo/trivy/pkg/iac/providers/github"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}
