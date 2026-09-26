package examples

import (
	"context"
	"fmt"
	"game-server/db"
	"game-server/db/wallet"

	"github.com/gofrs/uuid/v5"
)

func ExampleWallet() {
	e := db.NewExecutor(context.Background())

	user1 := uuid.UUID{}
	user2 := uuid.UUID{}

	// update the values for both user in a single tx
	// by default the Update uses the amount as a relative change so this means
	// add 15 coins to user1
	err := e.Atomic(context.Background(),
		wallet.UpdateCurrency(user1, "coins", wallet.Inc(15)),
		wallet.UpdateCurrency(user2, "gems", wallet.Inc(5)),
	)
	fmt.Printf("err: %v\n", err)

	// for removing currency just invert the sign
	err = e.Atomic(context.Background(),
		wallet.UpdateCurrency(user1, "coins", wallet.Dec(15)),
		wallet.UpdateCurrency(user2, "gems", wallet.Dec(5)),
	)
	fmt.Printf("err: %v\n", err)

	// and if you want to set the currency then
	err = e.Atomic(context.Background(),
		wallet.UpdateCurrency(user1, "coins", wallet.Set(0)),
		wallet.UpdateCurrency(user2, "gems", wallet.Set(0)),
	)
	fmt.Printf("err: %v\n", err)

	// allows to get the updated currency after the change
	// which included the value prev the change so the delta can be calculated
	op := wallet.UpdateCurrency(user1, "tokens", wallet.Inc(40))
	err = e.Atomic(context.Background(), op)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("op.Updated(): %v\n", op.Result())

	// can specify a options
	// clamp allows to specify a max value for the currency for that operation, in case
	// the currency has move value than the limit it will only update it's value
	// up to max, the applied delta can be calculated from the wallet.Result
	op = wallet.UpdateCurrency(user2, "tokens", wallet.Inc(40), wallet.WithClamp(100))
	err = e.Atomic(context.Background(), op)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("op.Updated(): %v\n", op.Result())

	// can specify to write to a ledger
	err = e.Atomic(context.Background(),
		wallet.UpdateCurrency(user1, "cupon", wallet.Inc(1), wallet.WriteLedger(wallet.LedgerEntry{"feature": "festival"})),
		wallet.UpdateCurrency(user2, "cupon", wallet.Inc(1), wallet.WriteLedger(wallet.LedgerEntry{"feature": "luckyWheel"})),
	)
	fmt.Printf("err: %v\n", err)
}
