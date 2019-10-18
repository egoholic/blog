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
To create DB you need to clone: https://github.com/stoa-bd/blogging-schema and then run in its directory:

```$ ./bin/createdb.sh <DB-PREFIX, like: production, development, test-acceptance>```

### How to seed DB?
To create DB you need to clone: https://github.com/stoa-bd/blogging-schema and then run in its directory:

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


### HTML framework

Use ID for unique elements on the page.
Use class for not unique elements on the page.
Use ``bhv-`` prefix for setting event handlers and in-browser acceptance-tests.
Use ``bhv-<entity/datapoint-name>`` to define the entity/datapoint representing context.
Use ``bhv-<entity/datapoint-name>__<attribute/nested-element name>`` to present attributes and nested content.

Example:
```
<article id="bhv-publication">
  <h1 id="bhv-publication__title">Title</h1>
	<div id="bhv-publication__content">
    Some content.
	</div>
</article>
```

Use ``ns-`` prefix for namespacing.

Example:
```
<div id="main">
  <div id="ns-recent" class="micro-list">
    <h3>Recent Publications</h3>
    <ol>
    {{range .RecentPublications}}
      {{template "publication--li" .}}
    {{end}}
    </ol>
  </div>
  <div id="ns-top" class="micro-list">
    <h3>Popular Publications</h3>
    <ol>
    {{range .PopularPublications}}
      {{template "publication--li" .}}
    {{end}}
    </ol>
  </div>
</div>
```
