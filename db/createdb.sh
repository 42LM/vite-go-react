#!/usr/bin/env sh

sqlite3 -batch "$PWD/vite-go-react.db" < "$PWD/db/initdb.sql"
