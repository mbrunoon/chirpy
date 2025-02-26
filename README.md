# chirpy

# Resources

## SQLC

## Goose
Migrations Manager: https://github.com/pressly/goose 

Create Migration:
goose create migration_name sql -dir db/migrations/

Migration Up/Down:
goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up -dir db/schema/

# Imports

go get github.com/lib/pq
go get github.com/google/uuid
go get github.com/joho/godotenv


# Postgres

sudo service postgresql start