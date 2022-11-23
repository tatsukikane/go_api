package graph
//go:generate go run github.com/99designs/gqlgen generate
import "github.com/tatsukikane/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//状態を追跡する場所

type Resolver struct{
	todos []*model.Todo
	//likeepsodeのデータの置き場所を記述
	likeEpisides []*model.LikeEpisode
}
