package priceList

import (
	"context"
	"fmt"
	myredis "gotest2/redis"
	"log"
	"sort"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func SamplePriceList() {
	data := []float64{19.9, 14.27, 10.3, 13.45, 10.92, 1.98, 12.07, 24.15, 14.96, 13.05, 17.32, 13.45, 16.2, 9.29, 13.45, 16.28, 13.94, 19.45, 12.54, 14.77, 23.53, 12.41, 16.27, 15.48, 13.51, 13.45, 12.56, 10.28, 22.13, 0.45, 16.19, 6.75, 18.97, 13.59, 13.39, 7.77, 19.98, 14.62, 20.23, 8.75, 18.94, 23.68, 5.99, 12.89, 11.85, 17.63, 10.36, 13.99, 15.48, 13.93, 10.52, 20.07, 24.15, 14.42}
	r := myredis.GetRedisStore("order-redis:6379")
	defer func() {
		r.Db.Close()
	}()
	ob := InitOrderBook("dddd", r.Db)

	for _, d := range data {
		ob = ob.InsertNewPrice(d)
		log.Println(ob.PLS)
	}

}

type OrderBook struct {
	ID  string `json:"id"`
	PLS []string
	R   *redis.Client
}

func (ob OrderBook) fetchPLS() OrderBook {
	cmd := ob.R.LRange(context.Background(), ob.ID, 0, -1)
	pls := cmd.Val()
	ob.PLS = pls
	return ob
}

func (ob OrderBook) fetchPLSWithPipe(pipe redis.Pipeliner) redis.Pipeliner {
	pipe.LRange(context.Background(), ob.ID, 0, -1)

	return pipe
}

func (ob OrderBook) insertNewpriceToRedis(pipe redis.Pipeliner, idx int, newprice float64) redis.Pipeliner {
	if len(ob.PLS) == idx {
		pipe.RPush(context.Background(), ob.ID, newprice)
		return pipe
	}

	newpriceS := fmt.Sprintf("%v", newprice)
	if ob.PLS[idx] != newpriceS {
		if idx == 0 {
			pipe.LPush(context.Background(), ob.ID, newprice)

		} else {
			pipe.LInsertAfter(context.Background(), ob.ID, ob.PLS[idx-1], newprice)
		}
	}
	return pipe
}

// func (ob OrderBook) setPriceHash(pipe redis.Pipeliner, price float64) redis.Pipeliner {
// 	pk := fmt.Sprintf("%v", price)
// 	pipe.HSet(context.Background(), ob.ID+"Hash", pk, "1234")
// 	return pipe
// }

func (ob OrderBook) InsertNewPrice(newprice float64) OrderBook {
	idx := sort.Search(len(ob.PLS), func(i int) bool {
		pf, _ := strconv.ParseFloat(ob.PLS[i], 64)
		return pf >= newprice
	})

	pipe := ob.R.TxPipeline()
	pipe = ob.insertNewpriceToRedis(pipe, idx, newprice)
	//pipe = ob.setPriceHash(pipe, newprice)
	pipe = ob.fetchPLSWithPipe(pipe)
	ress, _ := pipe.Exec(context.Background())
	ob.PLS = ress[len(ress)-1].(*redis.StringSliceCmd).Val()

	return ob
}

func InitOrderBook(id string, r *redis.Client) OrderBook {
	o := OrderBook{
		ID: id,
		R:  r,
	}
	o = o.fetchPLS()
	return o
}
