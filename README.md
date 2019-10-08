# Blog project

## FAQ

### Requirements
We use chromedriver for acceptance (Cucumber (godog)) tests:
```$ brew install chromedriver```

We store data with PostgreSQL 11 RDBMS. For local development we recommend to install PostgreSQL on MacOS with Postgres.app installer: https://postgresapp.com

### TODO: Using with Docker
TODO

### Configuration
```
$ go run targets/web/main.go help

	Name: Help (of type: 'Command')

		Description: Presents help information.

		Key:        help



	Name: Web server port (of type: 'Integer')

		Description: The port which web server listens to, like: 3000.

		ENV Var:    PORT
		Flag:       -port
		Key:        port



	Name: DB Host name (of type: 'String')

		Description: Database connection host name, like: 'localhost'.

		ENV Var:    DBHOST
		Flag:       -dbhost
		Key:        dbhost



	Name: DB port (of type: 'Integer')

		Description: Database connection port, like: 5432.

		ENV Var:    DBPORT
		Flag:       -dbport
		Key:        dbport



	Name: DB User (of type: 'String')

		Description: Database connection user name, like: 'postgres'.

		ENV Var:    DBUSER
		Flag:       -dbuser
		Key:        dbuser



	Name: DB Password (of type: 'String')

		Description: Database connection password.

		ENV Var:    DBPWD
		Flag:       -dbpwd
		Key:        dbpwd



	Name: Database name (of type: 'String')

		Description: Database name, like: 'stoa_blogging_development'.

		ENV Var:    DBNAME
		Flag:       -dbname
		Key:        dbname
```

### How to create DB and schema?

```$ ./bin/createdb.sh <DB-PREFIX, like: production, development, test-acceptance>```

### How to seed DB?

```$ ./bin/seed.sh <DB-PREFIX, like: production, development, test-acceptance>```

### How to run web-server?

```./bin/run.sh <DB-PREFIX, like: production, development, test-acceptance>```

It uses :3000 port.

### How to build?

```./bin/build.sh```


### How to run acceptance tests?

We use godog (Cucumber for Golang) and Agouti for acceptance tests via GUI.

You can run tests with:

```$ ./bin/test-acceptance.sh```
