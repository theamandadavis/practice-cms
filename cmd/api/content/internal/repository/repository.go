package repository

import (
	"context"

	content "github.com/theamandadavis/practice-cms/cmd/content/internal"

	"github.com/theamandadavis/practice-cms/pkg/id"

	"github.com/jmoiron/sqlx"
)

const (
	selectContent = `SELECT * FROM content WHERE id = ?`
)

type ContentRepo struct {
	DB *sqlx.DB
}

// New creates an instance of the contentRepo.
func New(conn *sqlx.DB) content.Repo {
	return &ContentRepo{conn}
}

// Get retrieves the article with the given id
func (r *ContentRepo) GetContentSummary(ctx context.Context, contentID id.UUID) (content.ContentSummary, error) {
	var contentSummary struct {
		ID           id.UUID `db:"id"`
		Title        string  `db:"title"`
		Body         string  `db:"body"`
		CreatedAt    string  `db:"created_at"`
		UpdatedAt    string  `db:"updated_at"`
		DeprecatedAt string  `db:"deprecated_at"`
	}

	if err := r.DB.GetContext(ctx, &contentSummary, selectContent, contentID); err != nil {
		return content.ContentSummary{}, err
	}

	return content.ContentSummary{
		ID:           contentSummary.ID,
		Title:        contentSummary.Title,
		Body:         contentSummary.Body,
		CreatedAt:    contentSummary.CreatedAt,
		UpdatedAt:    contentSummary.UpdatedAt,
		DeprecatedAt: contentSummary.DeprecatedAt,
	}, nil
}
