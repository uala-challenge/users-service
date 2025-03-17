package add_follow

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
)

type service struct {
	client *redis.Client
	log    log.Service
}

var _ Service = (*service)(nil)

func NewService(d Dependencies) *service {
	return &service{
		client: d.Client,
		log:    d.Log,
	}
}

func (s service) Accept(ctx context.Context, userID, followerID string) error {
	var err error
	maxRetries := 2
	retryDelay := 100 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		_, err = s.client.ZScore(ctx, "following:"+userID, followerID).Result()
		if err != redis.Nil {
			return s.log.WrapError(err, "Error verificando si el seguidor ya está en la lista")
		}

		_, err = s.client.ZAdd(ctx, "following:"+userID, redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: followerID,
		}).Result()
		if err == nil {
			break
		}

		_, err = s.client.ZAdd(ctx, "followers:"+followerID, redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: userID,
		}).Result()
		if err == nil {
			break
		}
		s.log.Info(ctx, "Reintentando...", map[string]interface{}{"Reintento": i + 1, "Esperando": retryDelay})
		time.Sleep(retryDelay)
	}
	if err != nil {
		return s.log.WrapError(err, "Error después de reintentos agregando seguidor")
	}

	return nil
}
