package repos

import (
	"database/sql"
	"fmt"

	"github.com/KL-Lru/sample-web-service/pkg/env"
	_ "github.com/lib/pq"
)

/*
	Create Connection Pool
*/
func Connect() (*sql.DB, error) {
	return sql.Open(
		"postgres",
		LoadConnectionInformation(),
	)
}

/*
	Load Environment variables
*/
func LoadConnectionInformation() string {
	user := env.GetEnvVal("POSTGRES_USER", "")
	password := env.GetEnvVal("POSTGRES_PASSWORD", "")
	host := env.GetEnvVal("POSTGRES_DB_HOST", "")
	dbname := env.GetEnvVal("POSTGRES_DATABASE", "")

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbname)
}

func (r *Repository) Ping() error {
	return r.Db.Ping()
}
