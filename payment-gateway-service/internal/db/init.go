package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"log"
)

func InitPostgresDB(cfg *config.Config) *sql.DB {
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
		log.Fatalln("Error when open database with exception : " + err.Error())
	}

	if err := con.Ping(); err != nil {
		log.Fatalln("Fail to ping the database please check it!!!")
	}

	return con
}
