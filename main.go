package main

import (
	"fmt"
	userService "github/Bendomey/peerstronix-store/user/cmd/user"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "domey"
	password = "akankobateng1"
	dbname   = "peerstronix"
)

func main() {
	userService.StartAccountService(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
}
