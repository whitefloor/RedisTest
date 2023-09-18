package main

// var ctx = context.Background()

type aaa struct {
	i int
}

func main() {
	// rdb := redisdb.InitRedisClient()

	// err := rdb.Set(ctx, "key", "1111", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// log.Println(rdb.Del(ctx, "key").Err())

	// val2, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("key", val2)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// res, err := rdb.ZRange(ctx, "123", 0, -1).Result()
	// log.Println(res, err)
}
