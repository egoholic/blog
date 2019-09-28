package schema

import (
	"database/sql"
	"fmt"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

func Apply(db *sql.DB) (err error) {
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", DBName))
	if err != nil {
		return
	}
	fmt.Printf("-- database `%s` has been dropped\n", DBName)
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", DBName))
	if err != nil {
		return
	}
	fmt.Printf("-- database `%s` has been created\n", DBName)

	db, err = sql.Open("postgres", DBConnectionString)
	defer db.Close()
	if err != nil {
		return
	}

	query := `CREATE TABLE rubrics (
		  slug             varchar(255) PRIMARY KEY,
			meta_keywords    text NOT NULL,
	    meta_description text NOT NULL,
	    title            varchar(255) NOT NULL,
	    description      text NOT NULL
		);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- table `rubrics` has been created")

	query = `CREATE TABLE publications (
		  slug             varchar(255) PRIMARY KEY,
			meta_keywords    text NOT NULL,
	    meta_description text NOT NULL,
	    title            varchar(255) NOT NULL,
	    content          text NOT NULL,
			created_at       timestamp without time zone NOT NULL,
			rubric_slug      varchar(255) NOT NULL REFERENCES rubrics(slug),
			popularity       int NOT NULL DEFAULT 0
		);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- table `publications` has been created")

	query = `CREATE INDEX publications_rubric_idx ON publications(rubric_slug);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- index `publications.rubric_slug` has been created")

	query = `CREATE TABLE accounts (
			login      varchar(255) PRIMARY KEY,
			first_name varchar(255) NOT NULL,
			last_name  varchar(255) NOT NULL,
	    bio        text NOT NULL
		);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- table `accounts` has been created")

	query = `CREATE TABLE publication_authors (
			publication_slug  varchar(255) NOT NULL REFERENCES publications(slug),
			author_login      varchar(255) NOT NULL REFERENCES accounts(login),

			PRIMARY KEY (publication_slug, author_login)
 		);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- table `publication_authors` has been created")

	query = `CREATE INDEX publication_author_login_idx ON publication_authors(author_login);`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	fmt.Println("-- index `accounts.publication_author` has been created")

	return
}
