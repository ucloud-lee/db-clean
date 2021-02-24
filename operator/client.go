package operator

import (
	"database/sql"
	"fmt"

	log "github.com/gogap/logrus"

	// to use mysql db
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	username = "app"
	passward = "bHa2xcizZ2HqvkNy1rNn4TSwdvmJIiobGEZQX3jH2R6OqbDvKEwByRrFEN01xUKH479CFqfrQXp7"
	network  = "tcp"
	host     = "[2002:ac1f:91c6::1]"
	port     = 10532
	database = "statement"
)

func (c *Clean) clientMysql() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, passward, network, host, port, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
