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
	_, err := s.client.ZScore(ctx, "following:"+followerID, userID).Result()
	if err == nil {
		return nil
	}
	if err != redis.Nil {
		return s.log.WrapError(err, "Error verificando si el seguidor ya está en la lista")
	}

	_, err = s.client.ZAdd(ctx, "following:"+followerID, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: userID,
	}).Result()
	if err != nil {
		return s.log.WrapError(err, "Error agregando seguidor en following")
	}
	_, err = s.client.ZAdd(ctx, "followers:"+userID, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: followerID,
	}).Result()
	if err != nil {
		return s.log.WrapError(err, "Error agregando seguidor en followers")
	}

	return nil
}
