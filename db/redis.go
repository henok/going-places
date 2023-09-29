package db

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}

// addToSet adds a value to a Redis set.
func AddToSet(key, value string) error {
    _, err := rdb.SAdd(ctx, key, value).Result()
    return err
}

// popRandomFromSet pops (removes) a random element from a Redis set.
func PopRandomFromSet(key string) (string, error) {
    // Get a random member from the set
    values, err := GetRandomMembersFromSet(key, 1)
    if err != nil {
        return "", err
    }
    if len(values) == 0 {
        return "", fmt.Errorf("set is empty")
    }
    randomMember := values[0]

    // Remove the random member from the set
    _, err = rdb.SRem(ctx, key, randomMember).Result()
    if err != nil {
        return "", err
    }

    return randomMember, nil
}

// getRandomMembersFromSet retrieves a specified number of random members from a Redis set.
func GetRandomMembersFromSet(key string, count int64) ([]string, error) {
    values, err := rdb.SRandMemberN(ctx, key, count).Result()
    return values, err
}
