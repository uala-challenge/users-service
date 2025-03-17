package add_follower

import (
	"context"

	"github.com/uala-challenge/users-service/internal/platform/remove_follow"

	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
	"github.com/uala-challenge/users-service/internal/platform/add_follow"
	"github.com/uala-challenge/users-service/internal/platform/update_timeline"
)

type Service interface {
	Accept(ctx context.Context, followed, follower string) error
}

type Dependencies struct {
	UpdateFollows  add_follow.Service
	RemoveFollows  remove_follow.Service
	UpdateTimeline update_timeline.Service
	Log            log.Service
}
