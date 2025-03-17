package update_timeline

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/uala-challenge/simple-toolkit/pkg/client/rest"
	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
)

type service struct {
	client rest.Service
	log    log.Service
}

var _ Service = (*service)(nil)

func NewService(d Dependencies) *service {
	return &service{
		client: d.Client,
		log:    d.Log,
	}
}

func (s service) Apply(ctx context.Context, user, follower string) (*resty.Response, error) {
	headers := map[string]string{
		"Aapply": "application/json",
	}
	body := map[string]string{
		"follower_id": follower,
	}

	rqt, err := json.Marshal(body)
	if err != nil {
		return nil, s.log.WrapError(err, "Error marshalling body")
	}

	rsp, err := s.client.Patch(ctx, fmt.Sprintf("/timeline/%s", user), rqt, headers)
	if err != nil {
		s.log.Error(ctx, err, "Error getting value", nil)
		return nil, err
	}
	return rsp, nil
}
