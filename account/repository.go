package account

import (
	"database/sql"
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
		return nil, err
	}
	return &postgresqlRepository{db}, nil
}

func (r *postgresqlRepository) Close() {
	// close db
	r.db.Close()
}
