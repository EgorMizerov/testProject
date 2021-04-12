package repository

import (
	"context"
	"fmt"
	"github.com/EgorMizerov/testProject/domain"
	"github.com/go-redis/redis/v8"
)

type HackerRedis struct {
	rdb *redis.Client
}

func NewHackerRedis(rdb *redis.Client) *HackerRedis {
	return &HackerRedis{rdb: rdb}
}

var hackers = map[string]float64{
	"Richard Stallman":   1953,
	"Alan Kay":           1940,
	"Yukihiro Matsumoto": 1965,
	"Claude Shannon":     1916,
	"Linus Torvalds":     1969,
}

func (r *HackerRedis) TestData() error {
	for key, value := range hackers {
		err := r.rdb.ZAdd(context.TODO(), "hackers", &redis.Z{
			Score:  value,
			Member: key,
		}).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *HackerRedis) GetHackers() ([]domain.Hacker, error) {
	res, err := r.rdb.ZRangeWithScores(context.TODO(), "hackers", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var hackers []domain.Hacker

	for _, i := range res {
		var hacker = domain.Hacker{
			fmt.Sprintf("%v", i.Member),
			i.Score,
		}
		hackers = append(hackers, hacker)
	}

	return hackers, nil
}
