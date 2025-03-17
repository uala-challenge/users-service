package patch_follow

import (
	"net/http"

	"github.com/uala-challenge/users-service/internal/add_follower"
)

type Service interface {
	Init(w http.ResponseWriter, r *http.Request)
}

type Dependencies struct {
	AddFollower add_follower.Service
}
