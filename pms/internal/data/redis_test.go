package data

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestHsetMap(t *testing.T) {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: []string{"127.0.0.1:6379"}})

	//err := rdb.HSet(context.Background(), "order:1", "1", Val{1: 1}).Err()
	//if err != nil {
	//	t.Error(err)
	//}
	//var val Val
	//if err := rdb.HGet(context.Background(), "order:1", "1").Scan(&val); err != nil {
	//	t.Error(err)
	//}
	//t.Log(val)
	rdb.HSet(context.Background(), "order:1", "product_id", "sku_id", "number", "price")
}
