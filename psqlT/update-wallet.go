package psqlT

import (
	"context"
	"database/sql"
)

type UpdateWalletParam struct {
	Name   string
	Amount float64
}

const updateWalletProductInfo = `
UPDATE WALLET
SET amount = $2
where name = $1
`

func UpdateWallet(ctx context.Context, tx *sql.Tx, param UpdateWalletParam) (res sql.Result, err error) {
	res, err = tx.ExecContext(ctx, updateWalletProductInfo, param.Name, param.Amount)

	return res, err

}
