package user

import (
	"database/sql"

	def "github.com/PhilSuslov/homework/iam/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

var _ def.IAMPostgresRepository = (*IAMRepo)(nil)

type IAMRepo struct {
	conn *pgxpool.Pool
}

func NewIAMRepo(conn *pgxpool.Pool) *IAMRepo {
	return &IAMRepo{conn: conn}
}

type Migrator struct {
	db            *sql.DB
	migrationsDir string
}

func NewMigrator(db *sql.DB, migrationsDir string) *Migrator {
	return &Migrator{
		db:            db,
		migrationsDir: migrationsDir,
	}
}

func (m *Migrator) Up() error {
	err := goose.Up(m.db, m.migrationsDir)
	if err != nil {
		return err
	}
	return nil
}
