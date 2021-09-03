package psqlT

import (
	"context"
	"database/sql"
)

type SelectWalletResult struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

const a = `
SELECT id, name, amount
FROM wallet
where name = $1
`

func SelectWallet(ctx context.Context, con *sql.Tx) (res SelectWalletResult, err error) {
	row := con.QueryRowContext(ctx, a, "tom")
	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.Amount,
	)
	return res, err
}
