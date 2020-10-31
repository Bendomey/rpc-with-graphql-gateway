package main

import (
	"fmt"
	"github/Bendomey/peerstronix-store/account/cmd/account"
)

const (
	host     = ""
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "peerstronix"
)

func main() {
	account.StartAccountService(fmt.Sprintf("host:%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, password, dbname, dbname))
}
