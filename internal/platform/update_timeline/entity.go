package update_timeline

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/uala-challenge/simple-toolkit/pkg/client/rest"
	"github.com/uala-challenge/simple-toolkit/pkg/utilities/log"
)

type Service interface {
	Apply(ctx context.Context, user, follower string) (*resty.Response, error)
}

type Dependencies struct {
	Client rest.Service
	Log    log.Service
}
