package user

import (
	"github/Bendomey/peerstronix-store/user"
	"log"
)

// Config holds the database url
type Config struct {
	databaseURL string
}

// StartUserService starts the user service
func StartUserService(url string) {
	var cfg = Config{
		databaseURL: url,
	}
	// get repository in here
	var r user.Repository
	r, err := user.NewPostgresqlRepository(cfg.databaseURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()
	//start server here
	log.Println("User starts on localhost 5000...")

	//start the grpc server

}
