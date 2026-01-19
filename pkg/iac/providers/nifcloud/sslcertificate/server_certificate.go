package sslcertificate

import (
	iacTypes "github.com/deliveroo/trivy/pkg/iac/types"
)

type ServerCertificate struct {
	Metadata   iacTypes.Metadata
	Expiration iacTypes.TimeValue
}
