package patch_follow

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/uala-challenge/users-service/internal/add_follower"

	"github.com/go-chi/chi/v5"
	"github.com/uala-challenge/users-service/kit"

	"github.com/uala-challenge/simple-toolkit/pkg/utilities/error_handler"
)

type service struct {
	upd add_follower.Service
}

var _ Service = (*service)(nil)

func NewService(d Dependencies) Service {
	return &service{
		upd: d.AddFollower,
	}
}

// Init godoc
// @Summary Update user timeline and add follower
// @Description Updates the timeline of a user when they are followed
// @Tags Timeline
// @Tags Follow
// @Accept json
// @Produce json
// @Param request body kit.Request true "query params"
// @Success 200 {object} map[string]string "Timeline updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /follow/{user_id} [patch]
func (s service) Init(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		_ = error_handler.HandleApiErrorResponse(error_handler.NewCommonApiError("bad request", err.Error(), err, http.StatusBadRequest), w)
		return
	}
	_ = r.Body.Close()

	rqt, _ := kit.BytesToModel[kit.Request](body)

	err = s.upd.Accept(r.Context(), userID, rqt.FollowerID)

	if err != nil {
		_ = error_handler.HandleApiErrorResponse(error_handler.NewCommonApiError("error fallowing", err.Error(), err, http.StatusInternalServerError), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Timeline updated successfully",
	})
	if err != nil {
		return
	}
}
