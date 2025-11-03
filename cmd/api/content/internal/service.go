package content

import (
	"context"

	"github.com/theamandadavis/practice-cms/pkg/id"
)

// Repo defines the DB level interaction of content
type Repo interface {
	GetContentSummary(ctx context.Context, id id.UUID) (ContentSummary, error)
}

// Service defines the service level contract that other services
// outside this package can use to interact with Article resources
type Service interface {
	GetContentSummary(ctx context.Context, id id.UUID) (ContentSummary, error)
}

type Content struct {
	repo Repo
}

// New Service instance
func New(repo Repo) Service {
	return &Content{repo}
}

// Get sends the request straight to the repo
func (s *Content) GetContentSummary(ctx context.Context, id id.UUID) (ContentSummary, error) {
	return s.repo.GetContentSummary(ctx, id)
}
