# Worklogger

Simple worklogger helper for daily end of day logging

## Requirements

- Go
- Postgres

## Uses

Create file inside DB folder and paste this code along with correct info

```go
package db

const (
    host     = ""
    port     = 0000
    user     = ""
    password = ""
    dbname   = ""
)
```

Run following commands

```sh
# Init database (only first time)
make

# Run application
make run
```

## License

[MIT](License)
