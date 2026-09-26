package db

import (
	"context"
	"slices"

	"github.com/jackc/pgx/v5"
)

type executor struct {
	conn *pgx.Conn
}

func NewExecutor(ctx context.Context) *executor {
	// TODO: implement
	panic("implement")
}

type Op interface {
	// This should be unique per operation
	OpKey() string
	// this adds the query to the batch
	PrepBatch(b *pgx.Batch)
	// this retrieves the results
	ProcessBatch(br pgx.BatchResults) error
}

// TODO: improve docs
// this function must run op in a single tx of the db (postgres)
// also should use batch to send a single round trip to the db
func (e *executor) Atomic(ctx context.Context, ops ...Op) error {
	sortedOps := make([]Op, 0, len(ops))
	indexedOps := make(map[string]int, len(ops))
	for i, op := range ops {
		sortedOps = append(sortedOps, op)
		indexedOps[op.OpKey()] = i
	}
	slices.SortFunc(sortedOps, func(a, b Op) int {
		if a.OpKey() < b.OpKey() {
			return -1
		} else if a.OpKey() > b.OpKey() {
			return 1
		} else {
			return 0
		}
	})

	var b pgx.Batch
	for _, op := range sortedOps {
		op.PrepBatch(&b)
	}

	tx, err := e.conn.BeginTx(ctx, pgx.TxOptions{
		// TBD: we could improve this
	})
	if err != nil {
		return err // TODO: can be improved
	}
	br := tx.SendBatch(ctx, &b)
	defer br.Close()
	for _, op := range sortedOps {
		err := op.ProcessBatch(br)
		if err != nil {
			// TBD: what to handle here
			// this represent errors reading the query from the db
			return err
		}
	}

	return nil
}
