package database

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/azure/database"
	"github.com/deliveroo/trivy/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) database.Database {
	return database.Database{
		MSSQLServers:      adaptMSSQLServers(modules),
		MariaDBServers:    adaptMariaDBServers(modules),
		MySQLServers:      adaptMySQLServers(modules),
		PostgreSQLServers: adaptPostgreSQLServers(modules),
	}
}
