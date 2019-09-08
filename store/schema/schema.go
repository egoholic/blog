package schema

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

func Apply(initCtx context.Context) (err error) {
	connStr, err := Config.DBCredentials().ConnectionStringWithoutDB()
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(initCtx, 10*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s;", Config.DBCredentials().DBName))
	if err != nil {
		return
	}
	fmt.Printf("-- database `%s` has been created\n", Config.DBCredentials().DBName)

	connStr, err = Config.DBCredentials().ConnectionString()
	if err != nil {
		return
	}
	db, err = sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return
	}
	query := `CREATE TABLE publications (
		  slug             varchar(255) PRIMARY KEY,
			meta_keywords    text NOT NULL,
	    meta_description text NOT NULL,
	    title            varchar(255),
	    content          text NOT NULL,
	    created_at       timestamp NOT NULL
		);`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return
	}
	fmt.Println("-- table `publications` has been created")

	query = `CREATE TABLE rubrics (
		  slug             varchar(255) PRIMARY KEY,
			meta_keywords    text NOT NULL,
	    meta_description text NOT NULL,
	    title            varchar(255),
	    description      text NOT NULL
		);`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return
	}
	fmt.Println("-- table `rubrics` has been created")
	return
}
