package all

import (
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/ansible"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/azurearm"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/cloudformation"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/dockerfile"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/helm"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/json"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/k8s"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/terraform"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/terraformplan/json"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/terraformplan/snapshot"
	_ "github.com/deliveroo/trivy/pkg/fanal/analyzer/config/yaml"
)
