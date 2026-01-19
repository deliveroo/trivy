package network

import (
	iacTypes "github.com/deliveroo/trivy/pkg/iac/types"
)

type VpnGateway struct {
	Metadata      iacTypes.Metadata
	SecurityGroup iacTypes.StringValue
}
