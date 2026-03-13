module github.com/PhilSuslov/homework/iam

go 1.25.2

replace github.com/PhilSuslov/homework/platform => ../platform

require (
	github.com/PhilSuslov/homework/platform v0.0.0-00010101000000-000000000000
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/gomodule/redigo v1.9.3
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.8.0
	github.com/pressly/goose/v3 v3.27.0
	github.com/samber/lo v1.53.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/text v0.34.0 // indirect
)
