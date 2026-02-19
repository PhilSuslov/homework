package order

import (
	"database/sql"

	def "github.com/PhilSuslov/homework/order/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
)

var _ def.OrderRepository = (*OrderRepo)(nil)

type OrderRepo struct {
	conn *pgx.Conn
}

func NewOrderRepo(conn *pgx.Conn) *OrderRepo {
	return &OrderRepo{conn: conn}

}

type Migrator struct {
	// mu     sync.RWMutex
	// orders map[string]*orderRepoModel.OrderDto
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
