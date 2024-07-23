package main

import (
	"context"
	"math/rand"
	"time"

	"fmt"
	"log"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	// Publish messages to different channels
	var randoms int
	var population, times string

	for {
		now := time.Now()
		times = now.Format("02-01-2006")
		randoms = rand.Int()
		population = fmt.Sprintf("Population: %d, time: %s", randoms, times)
		err := rdb.Publish(ctx, "Angliya", population).Err()
		if err != nil {
			log.Fatal(err)
		}

		now = time.Now()
		times = now.Format("02-01-2006")
		randoms = rand.Int()
		population = fmt.Sprintf("Population: %d, time: %s", randoms, times)
		err = rdb.Publish(ctx, "Ispaniya", population).Err()
		if err != nil {
			log.Fatal(err)
		}

		now = time.Now()
		times = now.Format("02-01-2006")
		randoms = rand.Int()
		population = fmt.Sprintf("Population: %d, time: %s", randoms, times)
		err = rdb.Publish(ctx, "Argentina", population).Err()
		if err != nil {
			log.Fatal(err)
		}

		now = time.Now()
		times = now.Format("02-01-2006")
		randoms = rand.Int()
		population = fmt.Sprintf("Population: %d, time: %s", randoms, times)
		err = rdb.Publish(ctx, "Brasiliya", population).Err()
		if err != nil {
			log.Fatal(err)
		}

		now = time.Now()
		times = now.Format("02-01-2006")
		randoms = rand.Int()
		population = fmt.Sprintf("Population: %d, time: %s", randoms, times)
		err = rdb.Publish(ctx, "India", population).Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Messages published successfully")
		time.Sleep(2 * time.Second)
	}
}
