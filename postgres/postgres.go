package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

// DBLogger struct defined
type DBLogger struct{}

// BeforeQuery function defined
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery function defined
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// New create a database connection to postgres
func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
