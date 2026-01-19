package seal

import (
	"slices"

	"github.com/deliveroo/trivy/pkg/detector/ospkg/driver"
	ftypes "github.com/deliveroo/trivy/pkg/fanal/types"
	"github.com/deliveroo/trivy/pkg/set"
)

var (
	supportedOSFamilies = set.New(
		ftypes.Alpine,
		ftypes.CBLMariner,
		ftypes.CentOS,
		ftypes.RedHat,
		ftypes.Debian,
		ftypes.Oracle,
		ftypes.Ubuntu,
	)
)

// Provider creates a Root.io driver if Root.io packages are detected
func Provider(osFamily ftypes.OSType, pkgs []ftypes.Package) driver.Driver {
	if supportedOSFamilies.Contains(osFamily) && slices.ContainsFunc(pkgs, sealPkg) {
		return NewScanner(osFamily)
	}
	return nil
}
