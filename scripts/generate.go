package scripts

//go:generate go run -mod=mod ../pkg/ent/entc.go
//go:generate go run -mod=mod ../internal/db/make_migrations.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
