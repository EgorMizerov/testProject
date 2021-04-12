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

func (r *HackerRedis) TestData(testTable map[string]float64) error {
	for key, value := range testTable {
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
