package main

import (
	_ "context"
	"fmt"
	_ "fmt"
	"github.com/go-redis/redis"
	gorma "github.com/hhxsv5/go-redis-memory-analysis"
)

var client *redis.Client

func redisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
		PoolSize: 20,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("redis init successfully")
	}
}

func writeAndAnalysis() {
	write(10000, "size10_10k", generateValue(10))
	write(20000, "size10_20k", generateValue(10))
	write(50000, "size10_50k", generateValue(10))

	write(10000, "size1000_10k", generateValue(1000))
	write(20000, "size1000_20k", generateValue(1000))
	write(50000, "size1000_50k", generateValue(1000))

	write(10000, "size5000_10k", generateValue(5000))
	write(20000, "size5000_20k", generateValue(5000))
	write(50000, "size5000_50k", generateValue(5000))

	analysis()
}

func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		_, err := client.Set(k, value, 0).Result()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection("127.0.0.1", 6379, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}

func main() {
	redisClient()
	writeAndAnalysis()
}
