package wallet

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type Result struct {
	PrevValue    int64
	CurrentValue int64
}

type currencyAmount struct {
	val uint64
	op  int8 // -1(dec)|0(set)|1(inc)
}

func Inc(v uint64) currencyAmount
func Dec(v uint64) currencyAmount
func Set(v uint64) currencyAmount

type updateOp struct{}

func (*updateOp) OpKey() string
func (*updateOp) PrepBatch(b *pgx.Batch)
func (*updateOp) ProcessBatch(br pgx.BatchResults) error
func (*updateOp) Result() Result

func UpdateCurrency(userId uuid.UUID, currencyType string, amount currencyAmount, opts ...opOpts) *updateOp
