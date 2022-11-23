package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tatsukikane/gqlgen-todos/graph/generated"
	"github.com/tatsukikane/gqlgen-todos/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
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
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		// CreatedAt: time.Now(),
		UserID:    input.UserID,
		EpisodeID: input.EpisodeID,
	}

	log.Println(likeEpisode)
	log.Printf("%T型", input.EpisodeID)

	r.likeEpisides = append(r.likeEpisides, likeEpisode)
	return likeEpisode, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// LikeEpisodes is the resolver for the likeEpisodes field.
func (r *queryResolver) LikeEpisodes(ctx context.Context) ([]*model.LikeEpisode, error) {
	return r.likeEpisides, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
