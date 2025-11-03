package content

import "practice-cms/pkg/id"

type ContentSummary struct {
	ID           id.UUID `db:"id"`
	Title        string  `db:"title"`
	Body         string  `db:"body"`
	CreatedAt    string  `db:"created_at"`
	UpdatedAt    string  `db:"updated_at"`
	DeprecatedAt string  `db:"deprecated_at"`
}
