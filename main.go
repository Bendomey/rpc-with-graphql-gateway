package main

import (
	"fmt"
	"github/Bendomey/peerstronix-store/account/cmd/account"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "domey"
	password = "akankobateng1"
	dbname   = "peerstronix"
)

func main() {
	account.StartAccountService(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
}
