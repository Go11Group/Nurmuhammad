package redis

import (
	"context"
	dbcon "new/storage/postgres"
	pb "new/structs"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

func GetUserById(id string, Repo *dbcon.UserRepo) (*pb.User, error) {
	rdb := ConnectDB()

	users, err := Repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userID := users.UserID
	user := map[string]interface{}{
		"user_id": users.UserID,
		"name":    users.Name,
		"email":   users.Email,
	}

	for k, v := range user {
		err := rdb.HSet(context.Background(), userID, k, v).Err()
		if err != nil {
			return nil, errors.Wrap(err, "failed to store user")
		}
	}

	expiration := 5 * time.Minute
	err = rdb.Expire(context.Background(), userID, expiration).Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to set expiration time")
	}

	userres, err := GetUsers(id)
	return userres, err
}

func GetUsers(id string) (*pb.User, error) {
	rdb := ConnectDB()

	userID := id
	order, err := rdb.HGetAll(context.Background(), userID).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve user data from Redis")
	}

	if len(order) == 0 {
		return nil, errors.New("no user found in Redis with the provided ID")
	}

	user := &pb.User{
		UserID: order["user_id"],
		Name:   order["name"],
		Email:  order["email"],
	}

	return user, nil
}
