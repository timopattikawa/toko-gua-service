package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timopattikawa/master-service-tokogua/cmd/config"
	"log"
)

func NewConnectionPsql(cfg *config.Config) *sql.DB {
	dns := fmt.Sprintf(
		"host=%s "+
			"port=%d "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable ",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name)

	con, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatalf("Fail to open database %s", err.Error())
		return nil
	}

	if err := con.Ping(); err != nil {
		panic(err.Error())
	}

	return con
}
