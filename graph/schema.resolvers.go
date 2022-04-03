package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/JeremyPersing/gographqltut/crud"
	"github.com/JeremyPersing/gographqltut/graph/generated"
	"github.com/JeremyPersing/gographqltut/graph/model"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	return crud.InsertDog(r.DB, input.Name, input.IsGoodBoi)
}

func (r *mutationResolver) UpdateDogName(ctx context.Context, id string, name string) (*model.Dog, error) {
	return crud.UpdateDogName(r.DB, id, name)
}

func (r *mutationResolver) DeleteDog(ctx context.Context, id string) (bool, error) {
	return crud.DeleteDog(r.DB, id)
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return crud.GetDog(r.DB, id)
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return crud.GetAllDogs(r.DB)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
