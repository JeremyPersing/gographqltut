package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/JeremyPersing/gographqltut/crud"
	"github.com/JeremyPersing/gographqltut/graph/generated"
	"github.com/JeremyPersing/gographqltut/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	return crud.InsertDog(r.DB, input.Name, input.IsGoodBoi)
}

func (r *mutationResolver) UpdateDogName(ctx context.Context, id string, name string) (*model.Dog, error) {
	var dog model.Dog

	res := r.DB.
		Model(&dog).
		// Model(&model.Dog{}). this works but returns empty data, w/out above Model called
		Clauses(clause.Returning{Columns: []clause.Column{}}).
		Where("id = ?", id).
		Update("name", name)

	return &dog, res.Error
}

func (r *mutationResolver) DeleteDog(ctx context.Context, id string) (bool, error) {
	deleted := true

	res := r.DB.Delete(&model.Dog{}, "id = ?", id)

	if res.Error != nil {
		deleted = false
		panic("Error deleting dog")
	}

	return deleted, res.Error
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
