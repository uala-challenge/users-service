package main

import (
	"github.com/uala-challenge/users-service/cmd/api/patch_follow"
	"github.com/uala-challenge/users-service/internal/add_follower"
	"github.com/uala-challenge/users-service/internal/platform/add_follow"
	"github.com/uala-challenge/users-service/internal/platform/remove_follow"
	"github.com/uala-challenge/users-service/internal/platform/update_timeline"
)

type repositories struct {
	UpdateTimeline update_timeline.Service
	AddFollow      add_follow.Service
	RemoveFollow   remove_follow.Service
}

type useCases struct {
	AddFollower add_follower.Service
}

type handlers struct {
	PatchFollow patch_follow.Service
}
