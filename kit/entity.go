package kit

type Request struct {
	FollowerID string `json:"follower_id" validate:"required"`
}
