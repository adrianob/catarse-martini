package config

import (
	"catarse/app/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

func Db() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=goDB sslmode=disable")
	if err != nil {
		fmt.Println("error")
	}
	//checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")
	dbmap.CreateTablesIfNotExists()

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	//err = dbmap.CreateTablesIfNotExists()
	//checkErr(err, "Create tables failed")

	return dbmap
}
