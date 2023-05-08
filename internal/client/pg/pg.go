package pg

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	Name     string
	QueryRaw string
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type NamedExecer interface {
	ScanOne(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAll(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type PG interface {
	QueryExecer
	NamedExecer
	Pinger
	Close() error
}

type pg struct {
	pool *pgxpool.Pool
}

func (p *pg) ScanOne(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%s: query context: %w", q.Name, err)
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAll(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%s: query context: %w", q.Name, err)
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	return p.pool.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	return p.pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) Close() error {
	p.pool.Close()
	return nil
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}
