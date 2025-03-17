package remove_follow

import (
	"context"

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
	_, err := s.client.ZScore(ctx, "following:"+userID, followerID).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return s.log.WrapError(err, "Error verificando si el seguidor existe en la lista")
	}

	_, err = s.client.ZRem(ctx, "following:"+userID, followerID).Result()
	if err != nil {
		return s.log.WrapError(err, "Error eliminando seguidor en following")
	}

	_, err = s.client.ZRem(ctx, "followers:"+followerID, userID).Result()
	if err != nil {
		return s.log.WrapError(err, "Error eliminando seguidor en followers")
	}

	return nil
}
