# Blog project

## FAQ

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

```$ go run targets/migrator/main.go```

### How to seed DB?

```$ go run targets/seed/main.go```

### How to run web-server?

```$ go run targets/web/main.go```

It uses :3000 port.

### How to run acceptance tests?

We use godog (Cucumber for Golang) and Agouti for acceptance tests via GUI.

You can run tests with:

```$ godog --random```
