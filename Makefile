db:
	@psql -u postgres -c 'create database worklogger;'

run:
	@go run ./main.go 
