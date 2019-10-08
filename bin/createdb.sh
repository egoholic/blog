#!/bin/bash
DBNAME=stoa_blogging_$1 go run ./targets/migrator/main.go
