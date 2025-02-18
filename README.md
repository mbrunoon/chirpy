# chirpy

# Resources

Migrations Manager: https://github.com/pressly/goose 

Create Migration:
goose create migration_name sql -dir db/migrations/

Migration Up/Down:
goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up -dir db/migrations/
