package azure

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/appservice"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/authorization"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/compute"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/container"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/cosmosdb"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/database"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/datafactory"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/datalake"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/keyvault"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/monitor"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/network"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/securitycenter"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/storage"
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/synapse"
)

type Azure struct {
	AppService     appservice.AppService
	Authorization  authorization.Authorization
	Compute        compute.Compute
	Container      container.Container
	CosmosDB       cosmosdb.CosmosDB
	Database       database.Database
	DataFactory    datafactory.DataFactory
	DataLake       datalake.DataLake
	KeyVault       keyvault.KeyVault
	Monitor        monitor.Monitor
	Network        network.Network
	SecurityCenter securitycenter.SecurityCenter
	Storage        storage.Storage
	Synapse        synapse.Synapse
}
