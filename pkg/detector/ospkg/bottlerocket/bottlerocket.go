package bottlerocket

import (
	"context"

	osver "github.com/deliveroo/trivy/pkg/detector/ospkg/version"
	ftypes "github.com/deliveroo/trivy/pkg/fanal/types"
	"github.com/deliveroo/trivy/pkg/log"
	"github.com/deliveroo/trivy/pkg/types"
)

// Scanner implements the Bottlerocket scanner
type Scanner struct {
}

// NewScanner is the factory method for Scanner
func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Detect(ctx context.Context, _ string, _ *ftypes.Repository, _ []ftypes.Package) ([]types.DetectedVulnerability, error) {
	log.InfoContext(ctx, "Vulnerability detection of Bottlerocket packages is currently not supported.")
	return nil, nil
}

func (s *Scanner) IsSupportedVersion(ctx context.Context, osFamily ftypes.OSType, osVer string) bool {
	return osver.Supported(ctx, nil, osFamily, osver.Minor(osVer))
}
