package wallet

type _opOpts struct{}
type opOpts = func(*_opOpts)

func WithClamp(v int64) func(opts *_opOpts)
func WriteLedger(entry LedgerEntry) func(opts *_opOpts)
