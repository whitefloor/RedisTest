package controller

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// Operation 覆寫測試，可以直接覆寫同一個key的value
func Operation(ctx context.Context, rdsClient *redis.Client) {
	// key:first
	err := rdsClient.Set(ctx, "key", "first", 0).Err()
	if err != nil {
		log.Println(err)
	}
	val, err := rdsClient.Get(ctx, "key").Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("key", val)

	// key:second 覆寫
	// 覆寫不會有任何反應
	err = rdsClient.Set(ctx, "key", "second", 0).Err()
	if err != nil {
		log.Println(err)
	}
	val2, err := rdsClient.Get(ctx, "key").Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("key", val2)

	// key:second 還沒被刪除，測試刪除兩次
	// first delete
	// Delete 操作會 return 1/0 代表成功或失敗
	cmd := rdsClient.Del(ctx, "key")
	i, err := cmd.Result()
	log.Printf("first delete %v %v", i, err) // result 1 , err nil

	// second delete
	cmd = rdsClient.Del(ctx, "key")
	i, err = cmd.Result()
	log.Printf("second delete %v %v", i, err) // result 0 , err nil

	// 刪除後取值
	// Get 操作可以用 redis.Nil 判断
	sCmd := rdsClient.Get(ctx, "key")
	if sCmd.Err() == redis.Nil {
		log.Println("Will print this log If redis has no value of key")
	}
	log.Println(sCmd.Result())                               // result "" , err redis nil
	log.Println("The value after second delete", sCmd.Val()) // val ""
}
