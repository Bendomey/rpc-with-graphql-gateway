package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Repository services to export from this module
type Repository interface {
	Close() /// to close database
}

//type to hold postgresql db
type postgresqlRepository struct {
	db *sql.DB
}

// NewPostgresqlRepository takes in the url of db and returns a repository or an error
func NewPostgresqlRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	//ping db
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	return &postgresqlRepository{db}, nil
}

// Close is used to close the db connection
func (r *postgresqlRepository) Close() {
	// close db
	r.db.Close()
}

func (r *postgresqlRepository) Ping() {
	//ping db for reponse
	r.db.Ping()
}
