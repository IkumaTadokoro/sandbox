package services

import (
	"context"
	"gqlgen-training/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

type Services interface {
	UserService
	RepoService
	IssueService
}

type services struct {
	*userService
	*repoService
	*issueService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:  &userService{exec: exec},
		repoService:  &repoService{exec: exec},
		issueService: &issueService{exec: exec},
	}
}
