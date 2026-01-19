package nas

import (
	iacTypes "github.com/deliveroo/trivy/pkg/iac/types"
)

type NASInstance struct {
	Metadata  iacTypes.Metadata
	NetworkID iacTypes.StringValue
}
