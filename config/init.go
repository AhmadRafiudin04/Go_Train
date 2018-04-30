package config

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "db_prod_edu"
)

func ConnString() string {

	strCon := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return fmt.Sprintf(strCon)
}
