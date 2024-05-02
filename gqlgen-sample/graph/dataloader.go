package graph

import (
	"context"
	"errors"
	"gqlgen-training/graph/model"
	"gqlgen-training/graph/services"

	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	indexes := make(map[string]int, len(IDs))
	for i, id := range IDs {
		indexes[id] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)

	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if user != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexes[user.ID]] = rsl
	}
	return results
}
