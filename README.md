# Blog project

## FAQ

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
