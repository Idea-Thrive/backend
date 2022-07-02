package mysql

import (
	"context"
	"database/sql"
	"fmt"

	// mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// New function.
func New(cfg Config, log *zap.Logger) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", CreateDataSource(cfg.Host, cfg.User, cfg.Pass, cfg.Port, cfg.Name))
	if err != nil {
		log.Fatal("error while connecting to mysql",
			zap.String("error", err.Error()),
		)

		return db, err //nolint:wrapcheck
	}

	ctx, cancel := context.WithCancel(context.Background())
	err = db.PingContext(ctx)

	if err != nil {
		cancel()
		log.Fatal("error while pinging mysql",
			zap.String("error", err.Error()),
		)

		return db, err //nolint:wrapcheck
	}

	cancel()

	return db, err //nolint:wrapcheck
}

// CreateDataSource .
func CreateDataSource(host, user, pass, port, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
}
