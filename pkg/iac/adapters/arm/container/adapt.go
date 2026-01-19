package container

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/container"
	"github.com/deliveroo/trivy/pkg/iac/scanners/azure"
)

func Adapt(deployment azure.Deployment) container.Container {
	return container.Container{
		KubernetesClusters: adaptKubernetesClusters(deployment),
	}
}

func adaptKubernetesClusters(_ azure.Deployment) []container.KubernetesCluster {

	return nil
}
