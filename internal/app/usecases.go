package app

import (
	"github.com/ridwanakf/tms-backend/internal/app/config"
)

type Usecases struct {
}

// Inject dependency for usecase layer
func newUsecases(cfg config.Config, repos *Repos) (*Usecases, error) {
	return &Usecases{}, nil
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
