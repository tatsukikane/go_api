package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/tatsukikane/gqlgen-todos/graph/generated"
	"github.com/tatsukikane/gqlgen-todos/graph/model"
)

// EpisodeID is the resolver for the episode_id field.
func (r *likeEpisodeResolver) EpisodeID(ctx context.Context, obj *model.LikeEpisode) (string, error) {
	panic(fmt.Errorf("not implemented: EpisodeID - episode_id"))
}

// UserID is the resolver for the user_id field.
func (r *likeEpisodeResolver) UserID(ctx context.Context, obj *model.LikeEpisode) (string, error) {
	panic(fmt.Errorf("not implemented: UserID - user_id"))
}

// CreatedAt is the resolver for the created_at field.
func (r *likeEpisodeResolver) CreatedAt(ctx context.Context, obj *model.LikeEpisode) (int, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// TODO: CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// LikeEpisodeを追加する
func (r *mutationResolver) AddLikeEpisode(ctx context.Context, input model.NewLikeEpisode) (*model.LikeEpisode, error) {
	likeEpisode := &model.LikeEpisode{
		ID:        int64(rand.Int()),
		CreatedAt: time.Now(),
		UserID:    input.UserID,
		EpisodeID: input.EpisodeID,
	}
	r.likeEpisides = append(r.likeEpisides, likeEpisode)
	return likeEpisode, nil
}

// TODO: Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// LikeEpisodes is the resolver for the likeEpisodes field.
func (r *queryResolver) LikeEpisodes(ctx context.Context) ([]*model.LikeEpisode, error) {
	panic(fmt.Errorf("not implemented: LikeEpisodes - likeEpisodes"))
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// LikeEpisode returns generated.LikeEpisodeResolver implementation.
func (r *Resolver) LikeEpisode() generated.LikeEpisodeResolver { return &likeEpisodeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type likeEpisodeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) LikeEpisode(ctx context.Context) ([]*model.LikeEpisode, error) {
	return r.likeEpisides, nil
}
