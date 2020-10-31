package account

import (
	"github/Bendomey/peerstronix-store/account"
	"log"
)

// Config holds the database url
type Config struct {
	databaseURL string
}

// StartAccountService starts the account service
func StartAccountService(url string) {
	var cfg = Config{
		databaseURL: url,
	}
	// get repository in here
	var r account.Repository
	r, err := account.NewPostgresqlRepository(cfg.databaseURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()
	//start server here
	log.Println("Account starts on localhost 5000...")

}
