package services

import (
	"context"
	"gqlgen-training/graph/db"
	"gqlgen-training/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

func convertIssue(issue *db.Issue) *model.Issue {
	return &model.Issue{
		ID:         issue.ID,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Author:     &model.User{ID: issue.Author},
		Repository: &model.Repository{ID: issue.Repository},
	}
}

func (i *issueService) GetIssueByNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssue(issue), nil
}
