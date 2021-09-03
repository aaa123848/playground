package main

import (
	"context"
	"database/sql"
	"gotest2/psqlT"
	myredis "gotest2/redis"
	"log"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

func Sum(a, b int) int {
	c := a + b
	return c
}

func main() {
	conn, err := psqlT.SqlConn()
	if err != nil {
		log.Println(err)
	}
	defer func() {
		conn.Close()
	}()
	r := myredis.GetRedisStore("order-redis:6379")
	defer func() {
		r.Db.Close()
	}()
	var wg sync.WaitGroup
	wg.Add(20)
	now := time.Now()
	for i := 0; i < 20; i++ {
		go Unit(&wg, conn, r)
	}
	wg.Wait()
	log.Println(time.Since(now))
}

func Unit(wg *sync.WaitGroup, conn *sql.DB, rstore myredis.RedisStore) {
	for i := 0; i < 100; i++ {
		ctx := context.TODO()
		lockId, err := GetLock(rstore)
		if err != nil {
			log.Println(err)
			continue
		}
		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
		updateWallet(ctx, tx, "tom", -0.01)
		rstore.UnLock(ctx, "iamlock", lockId)
	}
	wg.Done()
}

func GetLock(rstore myredis.RedisStore) (lockid string, err error) {
	for i := 0; i < 1000000; i++ {
		lockid, err = rstore.OnLock(context.TODO(), "iamlock")
		if err != nil {
			time.Sleep(1 * time.Microsecond)
			continue
		}
		break
	}
	return lockid, err
}

func updateWallet(ctx context.Context, tx *sql.Tx, name string, change float64) {
	tx.ExecContext(ctx, "SELECT amount FROM wallet WHERE name = $1 FOR UPDATE", name)
	res, err := psqlT.SelectWallet(ctx, tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}

	newAmount := decimal.NewFromFloat(res.Amount).Add(decimal.NewFromFloat(change))
	a, _ := newAmount.Float64()
	param := psqlT.UpdateWalletParam{
		Name:   name,
		Amount: a,
	}
	_, err = psqlT.UpdateWallet(ctx, tx, param)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	tx.Commit()

}
