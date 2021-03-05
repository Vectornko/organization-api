run:
	cmd/swag init -g cmd/main.go
	go run cmd/main.go

build:
	go build cmd/main.go

migrations:
	go run cmd/migrate.go --mode up

data:
	go run cmd/migrate.go --mode data

drop:
	go run cmd/migrate.go --mode drop