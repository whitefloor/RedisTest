package main

import (
	"RedisTest/controller"
	"RedisTest/db"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

func main() {
	rdsClient := db.InitRedisClient()
	// 操作測試
	controller.Operation(ctx, rdsClient)
	// Redis strings
	// string:string
	rdsClient.Set(ctx, "MyKey", "123", 60*time.Second)
	stringVal := rdsClient.Get(ctx, "MyKey").Val()
	log.Println("Get the key MyKey", stringVal) // result "123"
	// Redis Lists
	// 有序列表，基本上就stack結構的佇列
	rdsClient.LPush(ctx, "List", "1")
	rdsClient.LPush(ctx, "List", "2")
	listVal := rdsClient.LRange(ctx, "List", 0, -1).Val()
	log.Println("Lists value", listVal) // result [2 1]
	// Redis Sets
	// 無序的唯一集合，類似 hash map
	rdsClient.SAdd(ctx, "Set", "1", "2", "3")
	setVal := rdsClient.SMembers(ctx, "Set").Val()
	log.Println("Sets value", setVal)
	// Redis Sorted Sets
	// 有序的唯一集合，會依照 score 進行排列，相同的分數會依照 lexicographically 進行排序
	rdsClient.ZAdd(ctx, "ZSet",
		redis.Z{Score: 1, Member: "1"},
		redis.Z{Score: 1, Member: "2"},
		redis.Z{Score: 2, Member: "3"},
	)
	zsetVal := rdsClient.ZRange(ctx, "ZSet", 0, -1).Val()
	log.Println("ZSet value", zsetVal)
	// Redis Hashes
	// 具有結構化的集合，field-value
	rdsClient.HSet(ctx, "HSet", "1", "2", "3")
	hsetVal := rdsClient.HGetAll(ctx, "HSet").Val()
	log.Println(hsetVal)
}
