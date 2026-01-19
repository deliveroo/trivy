package google

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/google/bigquery"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/compute"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/dns"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/gke"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/iam"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/kms"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/sql"
	"github.com/deliveroo/trivy/pkg/iac/providers/google/storage"
)

type Google struct {
	BigQuery bigquery.BigQuery
	Compute  compute.Compute
	DNS      dns.DNS
	GKE      gke.GKE
	KMS      kms.KMS
	IAM      iam.IAM
	SQL      sql.SQL
	Storage  storage.Storage
}
