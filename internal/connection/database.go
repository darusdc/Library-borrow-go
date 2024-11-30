package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/darusdc/belajar-go/config"
	_ "github.com/lib/pq"
)

func GetDatabases(config config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		config.Host, config.Port, config.User, config.Pass, config.Name, config.Tz)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Error connection to DB", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Error to establish connection to DB", err.Error())
	}

	return db
}
