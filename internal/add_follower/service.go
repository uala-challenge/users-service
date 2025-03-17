package add_follower

import (
	"context"

	"github.com/uala-challenge/users-service/internal/platform/remove_follow"

	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
	"github.com/uala-challenge/users-service/internal/platform/add_follow"
	"github.com/uala-challenge/users-service/internal/platform/update_timeline"
)

type service struct {
	updateFollows  add_follow.Service
	updateTimeline update_timeline.Service
	removeFollows  remove_follow.Service
	log            log.Service
}

var _ Service = (*service)(nil)

func NewService(d Dependencies) *service {
	return &service{
		updateFollows:  d.UpdateFollows,
		updateTimeline: d.UpdateTimeline,
		removeFollows:  d.RemoveFollows,
		log:            d.Log,
	}
}

func (s service) Accept(ctx context.Context, userID, followerID string) error {
	err := s.updateFollows.Accept(ctx, userID, followerID)
	if err != nil {
		return s.log.WrapError(err, "Error actualizando follows")
	}
	_, err = s.updateTimeline.Apply(ctx, userID, followerID)
	if err == nil {
		return nil
	}
	err = s.removeFollows.Accept(ctx, userID, followerID)
	if err == nil {
		return nil
	}
	return s.log.WrapError(err, "Error eliminando follows")
}
