package content

import (
	"errors"
)

var (
	// ErrContentNotFound returns when the requested content is not found in the DB
	ErrContentNotFound = errors.New("requested content could not be found")
)
