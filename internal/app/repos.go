package app

import (
	"github.com/ridwanakf/tms-backend/internal/app/config"
)

type Repos struct {
}

// Inject dependency for repository layer
func newRepos(cfg config.Config) (*Repos, error) {
	return &Repos{}, nil
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
