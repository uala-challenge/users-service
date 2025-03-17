package main

import (
	"github.com/uala-challenge/simple-toolkit/pkg/simplify/app_builder"
	"github.com/uala-challenge/simple-toolkit/pkg/simplify/app_engine"
	"github.com/uala-challenge/users-service/cmd/api/patch_follow"
	"github.com/uala-challenge/users-service/internal/add_follower"
	"github.com/uala-challenge/users-service/internal/platform/add_follow"
	"github.com/uala-challenge/users-service/internal/platform/remove_follow"
	"github.com/uala-challenge/users-service/internal/platform/update_timeline"
)

type engine struct {
	simplify     app_engine.Engine
	repositories repositories
	useCases     useCases
	handlers     handlers
}

type AppBuilder struct {
	engine *engine
}

var _ app_builder.Builder = (*AppBuilder)(nil)

func NewAppBuilder() *AppBuilder {
	a := *app_engine.NewApp()
	return &AppBuilder{
		engine: &engine{
			simplify: a,
		},
	}
}

func (a engine) Run() error {
	return a.simplify.App.Run()
}

func (a AppBuilder) LoadConfig() app_builder.Builder {
	return a
}

func (a AppBuilder) InitRepositories() app_builder.Builder {
	a.engine.repositories.UpdateTimeline = update_timeline.NewService(update_timeline.Dependencies{
		Client: a.engine.simplify.RestClients["test"],
		Log:    a.engine.simplify.Log,
	})
	a.engine.repositories.AddFollow = add_follow.NewService(add_follow.Dependencies{
		Client: a.engine.simplify.RedisClient,
		Log:    a.engine.simplify.Log,
	})
	a.engine.repositories.RemoveFollow = remove_follow.NewService(remove_follow.Dependencies{
		Client: a.engine.simplify.RedisClient,
		Log:    a.engine.simplify.Log,
	})
	return a
}

func (a AppBuilder) InitUseCases() app_builder.Builder {
	a.engine.useCases.AddFollower = add_follower.NewService(add_follower.Dependencies{
		UpdateFollows:  a.engine.repositories.AddFollow,
		RemoveFollows:  a.engine.repositories.RemoveFollow,
		UpdateTimeline: a.engine.repositories.UpdateTimeline,
		Log:            a.engine.simplify.Log,
	})

	return a
}

func (a AppBuilder) InitHandlers() app_builder.Builder {
	a.engine.handlers.PatchFollow = patch_follow.NewService(patch_follow.Dependencies{
		AddFollower: a.engine.useCases.AddFollower})
	return a
}

func (a AppBuilder) InitRoutes() app_builder.Builder {
	a.engine.simplify.App.Router.Patch("/follow/{user_id}", a.engine.handlers.PatchFollow.Init)
	return a
}

func (a AppBuilder) Build() app_builder.App {
	return a.engine
}
