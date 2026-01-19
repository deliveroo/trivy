package ec2

import (
	iacTypes "github.com/deliveroo/trivy/pkg/iac/types"
)

type Subnet struct {
	Metadata            iacTypes.Metadata
	MapPublicIpOnLaunch iacTypes.BoolValue
}
