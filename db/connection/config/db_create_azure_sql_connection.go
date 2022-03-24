package config

import (
	connectionSql "github.com/aleinad25/db-sql-connection-gorm-go/db/connection"
)

type DatabaseConfiguration struct {
	DBServer   string
	DBPort     int
	DBUser     string
	DBName     string
	DBPassword string
}

func CreateAzureSQLConnection(configuration *DatabaseConfiguration) *connectionSql.DBConnection {
	azureSQLHost := configuration.DBServer
	azureSQLPort := configuration.DBPort
	azureSQLDatabase := configuration.DBName
	azureSQLUser := configuration.DBUser
	azureSQLPassword := configuration.DBPassword
	connection := connectionSql.NewSQLConnection(connectionSql.Config().
		SetSQLDialect(connectionSql.SQLServer).
		Host(azureSQLHost).
		Port(azureSQLPort).
		DatabaseName(azureSQLDatabase).
		User(azureSQLUser).
		Password(azureSQLPassword),
	)
	return connection
}
