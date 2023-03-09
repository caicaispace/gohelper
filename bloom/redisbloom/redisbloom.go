package redisbloom

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/willf/bloom"
)

func Run() {
	// 创建Redis客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 创建布隆过滤器
	bf := bloom.NewWithEstimates(1000000, 0.01)

	// 添加字符串到布隆过滤器中
	bf.AddString("apple")
	bf.AddString("banana")
	bf.AddString("pear")
	bf.AddString("peach")
	bf.AddString("grape")
	bf.AddString("watermelon")
	bf.AddString("strawberry")
	bf.AddString("cherry")
	bf.AddString("mango")
	// bf.AddString("orange")

	// 判断字符串是否存在于布隆过滤器中
	if bf.Test([]byte("apple")) {
		fmt.Println("The string 'apple' may exist")
	} else {
		fmt.Println("The string 'apple' doesn't exist")
	}

	if bf.Test([]byte("orange")) {
		fmt.Println("The string 'orange' may exist")
	} else {
		fmt.Println("The string 'orange' doesn't exist")
	}

	ctx := context.Background()
	// 将布隆过滤器保存到Redis中
	binary, err := bf.MarshalJSON()
	if err != nil {
		panic(err)
	}
	client.Set(ctx, "bloom_filter", binary, 0).Err()

	// 从Redis中读取布隆过滤器
	val, err := client.Get(ctx, "bloom_filter").Bytes()
	if err != nil {
		panic(err)
	}
	bf = &bloom.BloomFilter{}
	err = bf.UnmarshalJSON(val)
	if err != nil {
		panic(err)
	}

	// 判断字符串是否存在于Redis中的布隆过滤器中
	if bf.Test([]byte("apple")) {
		fmt.Println("The string 'apple' may exist in Redis bloom filter")
	} else {
		fmt.Println("The string 'apple' doesn't exist in Redis bloom filter")
	}

	if bf.Test([]byte("orange")) {
		fmt.Println("The string 'orange' may exist in Redis bloom filter")
	} else {
		fmt.Println("The string 'orange' doesn't exist in Redis bloom filter")
	}
}
