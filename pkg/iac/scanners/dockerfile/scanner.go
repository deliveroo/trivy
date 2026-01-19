package dockerfile

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/dockerfile"
	"github.com/deliveroo/trivy/pkg/iac/scanners/dockerfile/parser"
	"github.com/deliveroo/trivy/pkg/iac/scanners/generic"
	"github.com/deliveroo/trivy/pkg/iac/scanners/options"
	"github.com/deliveroo/trivy/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner[*dockerfile.Dockerfile] {
	defaultOpts := []options.ScannerOption{
		generic.WithSupportsInlineIgnore[*dockerfile.Dockerfile](true),
	}

	p := parser.NewParser()
	return generic.NewScanner("Dockerfile", types.SourceDockerfile, p, append(defaultOpts, opts...)...,
	)
}
