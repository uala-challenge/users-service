package remove_follow

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
)

type Service interface {
	Accept(ctx context.Context, followed, follower string) error
}

type Dependencies struct {
	Client *redis.Client
	Log    log.Service
}
