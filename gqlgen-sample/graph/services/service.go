package services

import (
	"context"
	"gqlgen-training/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByID(ctx context.Context, ID string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

type ProjectService interface {
	GetProjectV2ByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error)
	ListProjectByOwner(ctx context.Context, ownerID string, after *string, before *string, first *int, last *int) (*model.ProjectV2Connection, error)
}

type Services interface {
	UserService
	RepoService
	IssueService
	ProjectService
}

type services struct {
	*userService
	*repoService
	*issueService
	*projectService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:    &userService{exec: exec},
		repoService:    &repoService{exec: exec},
		issueService:   &issueService{exec: exec},
		projectService: &projectService{exec: exec},
	}
}
